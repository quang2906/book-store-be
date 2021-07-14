package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/model"
	repo "github.com/quang2906/book_store_be/repository"
)

func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	order := repo.GetAllOrder()
	responseWithJSON(w, http.StatusOK, order)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Invalid repository id"})
		return
	}
	products := repo.GetAllOrder()
	for _, repository := range products {
		if int(repository.Id) == id {
			responseWithJSON(w, http.StatusOK, repository)
			return
		}
	}
	responseWithJSON(w, http.StatusNotFound, map[string]string{"message": "Repository not found"})
}

func CreateOrder(writer http.ResponseWriter, request *http.Request) {
	var newOrder *model.Order
	if err := json.NewDecoder(request.Body).Decode(&newOrder); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	fmt.Println(newOrder)
	repo.CreateNewOrder(newOrder)
	responseWithJSON(writer, http.StatusCreated, newOrder)
}

func UpdateOrderById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid category id"})
		return
	}

	var updateOrder *model.Order
	if err := json.NewDecoder(request.Body).Decode(&updateOrder); err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid body"})
		return
	}
	updateOrder.Id = int64(id)

	orders := repo.GetAllProducts()
	for _, order := range orders {
		if order.Id == int64(id) {
			repo.UpdateOrderById(int64(id), updateOrder)
			responseWithJSON(writer, http.StatusOK, updateOrder)
			return
		}
	}

	responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Order not found"})
}

func DeleteOrderById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		responseWithJSON(writer, http.StatusBadRequest, map[string]string{"message": "Invalid Order id"})
		return
	}
	err = repo.DeleteOrderById(int64(id))
	if err != nil {
		responseWithJSON(writer, http.StatusNotFound, map[string]string{"message": "Order not found"})
		return
	}
	responseWithJSON(writer, http.StatusOK, map[string]string{"message": "Order was deleted"})
}
