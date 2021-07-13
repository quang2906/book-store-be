package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
