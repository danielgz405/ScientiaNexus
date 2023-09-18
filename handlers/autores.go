package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dg/acordia/models"
	"github.com/dg/acordia/repository"
	"github.com/dg/acordia/responses"
	"github.com/dg/acordia/server"
)

type InsertAutorRequest struct {
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Img       string `bson:"img" json:"img"`
	DesertRef string `bson:"desertRef" json:"desertRef"`
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
			responses.BadRequest(w, "Error creating user")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(profile)
	}
}
