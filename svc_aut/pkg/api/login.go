package api

import (
	"encoding/json"
	"github.com/d-Una-Interviews/svc_aut/pkg/model"
	. "github.com/d-Una-Interviews/svc_aut/pkg/security"
	"github.com/d-Una-Interviews/svc_aut/pkg/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	loginLogger, _ = zap.NewProduction(zap.Fields(zap.String("type", "user handler")))
	userService    service.UserService
)

func InitLogin(r *mux.Router) {
	userService = service.NewUserService()
	sr := r.PathPrefix("/user").Subrouter()
	sr.HandleFunc("/authenticate", authenticate).Methods("POST")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	loginRequest := &LoginRequest{}
	json.Unmarshal(reqBody, &loginRequest)
	var user model.User

	err2 := userService.FindUsername(r.Context(), &user, loginRequest.Username)

	if err2 != nil {
		var resp = map[string]interface{}{"status": false, "message": "User or password incorrect"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	//verificacion de la contraseña
	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))

	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	token, err := GenerateAccessToken(&user)
	if err != nil {
		loginLogger.Error(err.Error(), zap.Error(err))
	}

	var resp = map[string]interface{}{"message": "logged in"}
	resp["token"] = token //Store the token in the response
	json.NewEncoder(w).Encode(resp)
}
