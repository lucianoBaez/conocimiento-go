package utils

import (
	"encoding/json"
	"github.com/d-Una-Interviews/svc_aut/pkg/model"
	"net/http"
)

func PreconditionError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusPreconditionFailed)
	json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
	return
}

func BadRequestError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(model.Exception{Message: err.Error()})
	return
}

func CustomMessageError(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(model.Exception{Message: err})
	return
}
