package security

import (
	"context"
	"encoding/json"
	"github.com/d-Una-Interviews/svc_driver/internal/utils"
	"github.com/d-Una-Interviews/svc_driver/pkg/model"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Claims a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create the JWT key used to create the signature
var jwtKey = []byte("")

const (
	// Default JWT Env Var Name
	defaultJwtKeyEnvVar = "JWT_KEY"
	// Default JWT Env Var Value
	defaultJwtKeyValue = "my_secret_key"
)

func init() {
	jwtKey = []byte(utils.GetEnv(defaultJwtKeyEnvVar, defaultJwtKeyValue))
}

// JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(model.Exception{Message: "Missing auth token"})
			return
		}
		tk := &Claims{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
