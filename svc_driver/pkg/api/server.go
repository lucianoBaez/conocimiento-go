package api

import (
	"net/http"

	"github.com/d-Una-Interviews/svc_driver/internal/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	Router *mux.Router
}

var (
	Srv             *Server
	serverLogger, _ = zap.NewProduction(zap.Fields(zap.String("type", "repository")))
)

const (
	// Default port for REST Web API
	defaultRestPort = ":8081"
)

func NewServer() {
	serverLogger.Info("Server is initializing...")
	Srv = &Server{}
	Srv.Router = mux.NewRouter()
}

func StartServer() {
	var restPort = utils.GetEnv("REST_PORT", defaultRestPort)
	var handler http.Handler = Srv.Router

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	//originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(restPort, handlers.CORS(originsOk, headersOk, methodsOk)(handler))
}

func StopServer() {

}
