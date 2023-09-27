package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dg/acordia/models"
	"github.com/dg/acordia/repository"
	"github.com/dg/acordia/responses"
	"github.com/dg/acordia/server"
	"github.com/gorilla/mux"
)

type InsertAutorRequest struct {
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Img       string `bson:"img" json:"img"`
	DesertRef string `bson:"desertRef" json:"desertRef"`
}
type UpdateAutorRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Image     string `json:"image"`
	DesertRef string `json:"desertref"`
}

func InsertAutorHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req = InsertAutorRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}

		createAutor := models.InsertAutor{
			Email:     req.Email,
			Name:      req.Name,
			Img:       req.Img,
			DesertRef: req.DesertRef,
		}
		profile, err := repository.InsertAutor(r.Context(), &createAutor)
		if err != nil {
			responses.BadRequest(w, "Error creating Autor")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(profile)
	}
}

func UpdateAutorHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = UpdateAutorRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}
		data := models.UpdateAutor{
			Name:      req.Name,
			Email:     req.Email,
			Image:     req.Image,
			DesertRef: req.DesertRef,
		}
		params := mux.Vars(r)
		updatedAutor, err := repository.UpdateAutor(r.Context(), data, params["id"])
		if err != nil {
			fmt.Println(err)
			responses.BadRequest(w, "Error updating Autor")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedAutor)
	}
}

func DeleteAutorHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err := repository.DeleteAutor(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting user")
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func ListAutorHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		autor, err := repository.ListAutor(r.Context())
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		if autor == nil {
			autor = []models.Autor{}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(autor)
	}
}

func GetAutorByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		article, err := repository.GetAutorById(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(article)
	}
}
