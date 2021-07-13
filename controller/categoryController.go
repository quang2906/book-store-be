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

func CreateCategory(writer http.ResponseWriter, request *http.Request) {
	var newCategory *model.Category
	if err := json.NewDecoder(request.Body).Decode(&newCategory); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	repo.CreateCategory(newCategory)
	responseWithJSON(writer, http.StatusCreated, newCategory)
}

func UpdateCategory(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid category id"})
		return
	}

	var updateCategory *model.Category
	if err := json.NewDecoder(request.Body).Decode(&updateCategory); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateCategory.Id = int64(id)

	categories := repo.GetAllCategories()
	for _, category := range categories {
		if category.Id == int64(id) {
			repo.UpdateCategoryById(int64(id), updateCategory)
			responseWithJSON(writer, http.StatusOK, updateCategory)
			return
		}
	}

	responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Todo not found"})
}

func DeleteCategoryById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid category id"})
		return
	}

	categories := repo.GetAllCategories()
	for _, category := range categories {
		if category.Id == int64(id) {
			repo.DeleteCategoryById(int64(id))
			responseWithJSON(writer, http.StatusOK, map[string]string{"message": "Category was deleted"})
			return
		}
	}

	responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Category not found"})
}

func responseWithJSON(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(object)
}
