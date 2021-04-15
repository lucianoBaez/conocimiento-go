package api

import (
	"github.com/d-Una-Interviews/svc_driver/internal/utils"
	"github.com/d-Una-Interviews/svc_driver/pkg/security"
	"net/http"
)

const (
	// Default port for REST Web API
	pathPrefix = "/api/v1"
)

//InitAPI Initialize Web REST API
func InitAPI() {
	var (
		pathPrefix = utils.GetEnv("PATH_PREFIX", pathPrefix)
	)
	routerWithSecurity := Srv.Router.PathPrefix(pathPrefix).Subrouter()
	routerWithoutSecurity := Srv.Router.PathPrefix(pathPrefix).Subrouter()

	InitDriver(routerWithSecurity)

	routerWithoutSecurity.Use(commonMiddleware)
	routerWithSecurity.Use(security.JwtVerify)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
