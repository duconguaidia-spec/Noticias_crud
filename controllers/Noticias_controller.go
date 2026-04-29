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

func GetAllNoticias(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		"SELECT id, titulo, tipo, cuerpo, url_video, imagen_destacada, fuente, fecha_noticia, id_usuario, acceso_limitado, estado, activo, fecha_creacion, fecha_modificacion FROM noticias",
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()
	var list []models.Noticia
	for rows.Next() {
		var n models.Noticia
		if err := rows.Scan(&n.ID, &n.Titulo, &n.Tipo, &n.Cuerpo, &n.UrlVideo, &n.ImagenDestacada, &n.Fuente, &n.FechaNoticia, &n.IDUsuario, &n.AccesoLimitado, &n.Estado, &n.Activo, &n.FechaCreacion, &n.FechaModificacion); err != nil {
			respondJSON(w, 500, map[string]string{"error": err.Error()})
			return
		}
		list = append(list, n)
	}
	if err := rows.Err(); err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	respondJSON(w, 200, list)
}

func GetNoticiaByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var n models.Noticia
	err := config.DB.QueryRow(
		"SELECT id, titulo, tipo, cuerpo, url_video, imagen_destacada, fuente, fecha_noticia, id_usuario, acceso_limitado, estado, activo, fecha_creacion, fecha_modificacion FROM noticias WHERE id=$1", id,
	).Scan(&n.ID, &n.Titulo, &n.Tipo, &n.Cuerpo, &n.UrlVideo, &n.ImagenDestacada, &n.Fuente, &n.FechaNoticia, &n.IDUsuario, &n.AccesoLimitado, &n.Estado, &n.Activo, &n.FechaCreacion, &n.FechaModificacion)
	if err != nil {
		respondJSON(w, 404, map[string]string{"error": "Noticia no encontrada"})
		return
	}
	respondJSON(w, 200, n)
}

