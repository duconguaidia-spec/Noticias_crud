package routes

import (
	"API_GO_CRUD/controllers"
	"github.com/gorilla/mux"
)

func RegisterNoticiaRoutes(r *mux.Router) {
	r.HandleFunc("/noticias", controllers.GetAllNoticias).Methods("GET")
	r.HandleFunc("/noticias/{id}", controllers.GetNoticiaByID).Methods("GET")
	r.HandleFunc("/noticias", controllers.CreateNoticia).Methods("POST")
	r.HandleFunc("/noticias/{id}", controllers.UpdateNoticia).Methods("PUT")
	r.HandleFunc("/noticias/{id}", controllers.DeleteNoticia).Methods("DELETE")
}