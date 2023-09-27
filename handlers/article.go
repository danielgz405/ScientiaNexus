package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dg/acordia/models"
	"github.com/dg/acordia/repository"
	"github.com/dg/acordia/responses"
	"github.com/dg/acordia/server"
	"github.com/gorilla/mux"
)

type InsertArticleRequest struct {
	Name      string   `bson:"name" json:"name"`
	Autor     string   `bson:"autor" json:"autor"`
	Content   string   `bson:"content" json:"content"`
	Documents []string `bson:"documents" json:"documents"`
	Image     string   `bson:"image" json:"image"`
	DesertRef string   `bson:"desertRef" json:"desertRef"`
}

type UpdateArticleRequest struct {
	Name      string   `bson:"name" json:"name"`
	Autor     string   `bson:"autor" json:"autor"`
	Content   string   `bson:"content" json:"content"`
	Documents []string `bson:"documents" json:"documents"`
	Image     string   `bson:"image" json:"image"`
	DesertRef string   `bson:"desertRef" json:"desertRef"`
}

func InsertArticleHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = InsertArticleRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		article := models.InsertArticle{}
		currentTime := time.Now()
		formattedTime := currentTime.Format("2006-01-02 15:04:05")

		article = models.InsertArticle{
			Name:      req.Name,
			Autor:     req.Autor,
			Date:      formattedTime,
			Content:   req.Content,
			Documents: req.Documents,
			Image:     req.Image,
			DesertRef: req.DesertRef,
		}

		createdArticle, err := repository.InsertArticle(r.Context(), &article)
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdArticle)
	}
}

func ListArticlesHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		articles, err := repository.ListArticles(r.Context())
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		if articles == nil {
			articles = []models.Article{}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(articles)
	}
}

func UpdateArticleHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var req = UpdateArticleRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request")
			return
		}
		article := models.UpdateArticle{}

		article = models.UpdateArticle{
			Name:      req.Name,
			Autor:     req.Autor,
			Content:   req.Content,
			Documents: req.Documents,
			Image:     req.Image,
			DesertRef: req.DesertRef,
		}

		updatedArticle, err := repository.UpdateArticle(r.Context(), article, params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedArticle)
	}
}
func DeleteArticleHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err := repository.DeleteArticle(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		responses.DeleteResponse(w, "Article deleted")
	}
}

func GetArticleByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		article, err := repository.GetArticleById(r.Context(), params["id"])
		if err != nil {
			responses.InternalServerError(w, err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(article)
	}
}
