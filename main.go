package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dg/acordia/handlers"
	"github.com/dg/acordia/middleware"
	"github.com/dg/acordia/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DB_URI := os.Getenv("DB_URI")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      ":" + PORT,
		JWTSecret: JWT_SECRET,
		DbURI:     DB_URI,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.Use(middleware.CheckAuthMiddleware(s))
	r.HandleFunc("/welcome", handlers.HomeHandler(s)).Methods(http.MethodGet)

	//Auth
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	//user
	r.HandleFunc("/user/delete", handlers.DeleteUserHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/user/update", handlers.UpdateUserHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/user/profile/{userId}", handlers.ProfileHandler(s)).Methods(http.MethodGet)

	//autor
	r.HandleFunc("/autor", handlers.InsertAutorHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/autor/delete/{id}", handlers.DeleteAutorHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/autor/update/{id}", handlers.UpdateAutorHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/autor/one/{id}", handlers.GetAutorByIdHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/autor/list", handlers.ListAutorHandler(s)).Methods(http.MethodGet)

	//aticle
	r.HandleFunc("/article", handlers.InsertArticleHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/article/delete/{id}", handlers.DeleteArticleHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/article/update/{id}", handlers.UpdateArticleHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/article/list", handlers.ListArticlesHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/article/one/{id}", handlers.GetArticleByIdHandler(s)).Methods(http.MethodGet)

}
