package main

import (
	"log"
	"net/http"
	"user-service/handlers"
	"user-service/pkg/repository"
	"user-service/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := repository.NewDB()

	service := service.NewService(db)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	log.Println("localhost:8000")

	r.Post("/user", handlers.RegisterUser(service))
	r.Get("/users", handlers.ListUsers(service))
	r.Get("/user/{id}", handlers.GetUserBYID(service))

	log.Println(http.ListenAndServe("localhost:8000", r))

}
