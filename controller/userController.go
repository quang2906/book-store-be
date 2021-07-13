package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/model"
	repo "github.com/quang2906/book_store_be/repository"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := repo.GetAllUsers()
	responseWithJSON(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid user id"})
		return
	}

	users := repo.GetAllUsers()
	for _, repository := range users {
		if int(repository.Id) == id {
			responseWithJSON(w, http.StatusOK, repository)
			return
		}
	}

	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "User not found"})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser *model.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	repo.CreateUser(newUser)
	responseWithJSON(w, http.StatusCreated, newUser)
}
