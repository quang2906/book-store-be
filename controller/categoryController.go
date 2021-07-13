package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/model"
	repo "github.com/quang2906/book_store_be/repository"
)

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	categories := repo.GetAllCategories()
	responseWithJSON(w, http.StatusOK, categories)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid repository id"})
		return
	}

	categories := repo.GetAllCategories()
	for _, repository := range categories {
		if int(repository.Id) == id {
			responseWithJSON(w, http.StatusOK, repository)
			return
		}
	}

	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "Repository not found"})
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory *model.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	repo.CreateCategory(newCategory)
	responseWithJSON(w, http.StatusCreated, newCategory)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid category id"})
		return
	}

	var updateCategory *model.Category
	if err := json.NewDecoder(r.Body).Decode(&updateCategory); err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateCategory.Id = int64(id)

	categories := repo.GetAllCategories()
	for _, category := range categories {
		if category.Id == int64(id) {
			repo.UpdateCategoryById(int64(id), updateCategory)
			responseWithJSON(w, http.StatusOK, updateCategory)
			return
		}
	}

	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid category id"})
		return
	}

	categories := repo.GetAllCategories()
	for _, category := range categories {
		if category.Id == int64(id) {
			repo.DeleteCategoryById(int64(id))
			responseWithJSON(w, http.StatusOK, map[string]string{"message": "Category was deleted"})
			return
		}
	}

	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "Category not found"})
}

func responseWithJSON(w http.ResponseWriter, status int, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(object)
}
