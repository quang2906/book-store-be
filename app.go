package main

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/routers"
)

func main() {
	database.Connect()
	defer database.DB.Clauses()
	r := mux.NewRouter()
	routers.ConfigUserRouter(r)
	routers.ConfigRouter(r)
	// routers.ConfigFileRouter(r)
	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./public/")))
	r.PathPrefix("/images/").Handler(images)
	http.ListenAndServe(":3000", handlers.CORS()(r))

}
