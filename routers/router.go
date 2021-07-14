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

func ConfigProductsRouter(router *mux.Router) {
	product := router.PathPrefix("/products").Subrouter()
	product.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllProducts)
	product.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetProductById)
	product.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateProduct)
	product.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateProductById)
	product.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteProductById)
}
