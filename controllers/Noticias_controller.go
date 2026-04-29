package controllers

import (
	"API_GO_CRUD/config"
	"API_GO_CRUD/models"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
