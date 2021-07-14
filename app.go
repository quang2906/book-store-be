package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/routers"
)

func main() {
	database.Connect()
	defer database.DB.Clauses()
	r := mux.NewRouter()
	routers.ConfigCategoryRouter(r)
	routers.ConfigProductsRouter(r)
	routers.ConfigUserRouter(r)

	http.ListenAndServe(":3000", r)

}
