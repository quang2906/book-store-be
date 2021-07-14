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
	for _, user := range users {
		if int(user.Id) == id {
			responseWithJSON(w, http.StatusOK, user)
			return
		}
	}

	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "User not found"})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	newUser := model.User{
		Name:        data["name"],
		PhoneNumber: data["phone_number"],
		Email:       data["email"],
		Role:        data["role"],
	}
	newUser.HashPassword(data["password"])
	repo.CreateUser(&newUser)
	responseWithJSON(w, http.StatusCreated, newUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid user id"})
		return
	}

	var updateUser *model.User
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateUser.Id = int64(id)

	users := repo.GetAllUsers()
	for _, user := range users {
		if user.Id == int64(id) {
			repo.UpdateUserById(int64(id), updateUser)
			responseWithJSON(w, http.StatusOK, updateUser)
			return
		}
	}

	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "User not found"})
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid user id"})
		return
	}

	users := repo.GetAllUsers()
	for _, user := range users {
		if user.Id == int64(id) {
			repo.DeleteUserById(int64(id))
			responseWithJSON(w, http.StatusOK, map[string]string{"message": "user was deleted"})
			return
		}
	}

	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "User not found"})
}
