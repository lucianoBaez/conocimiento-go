package security

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/d-Una-Interviews/svc_aut/internal/utils"
	"github.com/d-Una-Interviews/svc_aut/pkg/model"

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

// GenerateAccessToken Create a JWT for the user
func GenerateAccessToken(user *model.User) (string, error) {

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}
	return tokenString, nil
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
