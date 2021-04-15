package api

import (
	"context"
	"encoding/json"
	"github.com/d-Una-Interviews/svc_driver/pkg/model"
	"github.com/d-Una-Interviews/svc_driver/pkg/repository"
	"github.com/d-Una-Interviews/svc_driver/pkg/service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

var (
	driverLogger, _  = zap.NewProduction(zap.Fields(zap.String("type", "user handler")))
	driverService    service.DriverService
	driverRepository = repository.InitMongoRepository()
)

func InitDriver(r *mux.Router) {
	driverService = service.NewdriverService(model.NewDriverRepository(driverRepository))
	sr := r.PathPrefix("/drivers").Subrouter()
	sr.HandleFunc("", FindAllDrivers).Methods("GET")
	sr.HandleFunc("/radius/{radius}", FindByRadius).Methods("GET")
	sr.HandleFunc("/limit/{limit}/page/{page}", FindAllPaged).Methods("GET")
	sr.HandleFunc("", createDriver).Methods("POST")
}

func FindByRadius(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	radius := vars["radius"]

	intValue, error := strconv.Atoi(radius)
	if error != nil {
		driverLogger.Error("Error getting int value")
	}

	err, result := driverService.FindByRadius(r.Context(), intValue)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Exception{Message: "drivers not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func createDriver(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var driver model.Driver
	json.Unmarshal(reqBody, &driver)

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	err := driverService.Create(ctx, &driver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Exception{Message: "Error creating driver"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(driver)
}

func FindAllDrivers(w http.ResponseWriter, r *http.Request) {
	err, result := driverService.FindAll(r.Context())
	if err != nil {
		driverLogger.Error(err.Error(), zap.Error(err))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func FindAllPaged(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	page := vars["page"]

	limitValue, error := strconv.Atoi(limit)
	pageValue, error := strconv.Atoi(page)
	if error != nil {
		driverLogger.Error("Error getting int value")
	}

	err, result := driverService.FindAllPaged(r.Context(), int64(limitValue), int64(pageValue))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.Exception{Message: "driver not found"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
