package api

import (
	"context"
	"encoding/json"
	"github.com/d-Una-Interviews/svc_aut/pkg/model"
	"github.com/d-Una-Interviews/svc_aut/pkg/repository"
	"github.com/d-Una-Interviews/svc_aut/pkg/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

var (
	userLogger, _  = zap.NewProduction(zap.Fields(zap.String("type", "user handler")))
	usrService     service.UserService
	userRepository = repository.InitRepository()
)

func InitUser(r *mux.Router) {
	usrService = service.NewUserService()
	sr := r.PathPrefix("/users").Subrouter()
	sr.HandleFunc("", FindAll).Methods("GET")
	sr.HandleFunc("/username/{username}", FindByUsername).Methods("GET")
	sr.HandleFunc("", create).Methods("POST")
}

func FindByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	var user model.User
	err := usrService.FindUsername(r.Context(), &user, username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Exception{Message: "User not found"})
		return
	}
	json.NewEncoder(w).Encode(user)
}

func create(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user model.User
	json.Unmarshal(reqBody, &user)

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	createUser, err := usrService.CreateUser(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Exception{Message: "Error creating user"})
		return
	}
	json.NewEncoder(w).Encode(createUser)
}

// ListUsers List users in the system
func FindAll(w http.ResponseWriter, r *http.Request) {
	var result []model.User
	err := usrService.FindAll(r.Context(), &result)
	if err != nil {
		userLogger.Error(err.Error(), zap.Error(err))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
