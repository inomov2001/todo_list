package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"user-service/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func RegisterUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// byteStructJSON := r.Body
		// JSON, err := io.ReadAll(byteStructJSON)
		// if err != nil{
		// 	render.JSON(w, r, map[string]bool{"error": true})
		// 	return
		// }

		// user := service.User{}		
		// json.NewDecoder(byteStructJSON).Decode(user)
		// res, err := svc.RegisterUser(context.Background(), user.Name)
		var u service.User
		if err := render.DecodeJSON(r.Body, &u); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}
		id, err := svc.RegisterUser(context.Background(), u.Name)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusCreated)
		render.JSON(w, r, render.M{
			"id": id,
		})
	}
}

func UpdateUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteUser(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func ListUsers(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := svc.ListUsers(context.Background())
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, users)
	}
}

func GetUserBYID(svc service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("1")
		userID := chi.URLParam(r, "id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, render.M{
				"error": fmt.Sprintf("path parameter 'id' isn't a number: %q", userID),
			})

		}
		user, err := svc.GetUserBYID(context.Background(), id)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, render.M{
				"error": err.Error(),
			})
			return
		}
		render.JSON(w, r, user)
	}
}