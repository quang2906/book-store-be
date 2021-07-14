package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/model"
	repo "github.com/quang2906/book_store_be/repository"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := repo.GetAllProducts()
	responseWithJSON(w, http.StatusOK, products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid repository id"})
		return
	}
	products := repo.GetAllProducts()
	for _, repository := range products {
		if int(repository.Id) == id {
			responseWithJSON(w, http.StatusOK, repository)
			return
		}
	}
	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "Repository not found"})
}

func CreateProduct(writer http.ResponseWriter, request *http.Request) {
	var newProduct *model.Product
	if err := json.NewDecoder(request.Body).Decode(&newProduct); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}

	repo.CreateNewProduct(newProduct)
	responseWithJSON(writer, http.StatusCreated, newProduct)
}

func UpdateProductById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid category id"})
		return
	}

	var updateProducts *model.Product
	if err := json.NewDecoder(request.Body).Decode(&updateProducts); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateProducts.Id = int64(id)

	products := repo.GetAllProducts()
	for _, product := range products {
		if product.Id == int64(id) {
			repo.UpdateProductById(int64(id), updateProducts)
			responseWithJSON(writer, http.StatusOK, updateProducts)
			return
		}
	}

	responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Product not found"})
}

func DeleteProductById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid Product id"})
		return
	}
	err = repo.DeleteProductById(int64(id))
	if err != nil {
		responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Product not found"})
		return
	}
	responseWithJSON(writer, http.StatusOK, map[string]string{"message": "Product was deleted"})
}
