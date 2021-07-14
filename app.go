package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/routers"
)

func main() {
	fmt.Println("Hello Tris")
	database.Connect()
	defer database.DB.Clauses()
	r := mux.NewRouter()
	routers.ConfigCategoryRouter(r)
	routers.ConfigProductsRouter(r)

	http.ListenAndServe(":3000", r)

}
