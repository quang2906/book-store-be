package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/quang2906/book_store_be/controller"
)

func ConfigCategoryRouter(router *mux.Router) {
	category := router.PathPrefix("/categories").Subrouter()
	category.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllCategory)
	category.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetCategoryById)
	category.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateCategory)
	category.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateCategory)
	category.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteCategoryById)

}
