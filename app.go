package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/routers"
	"github.com/rs/cors"
	"github.com/stripe/stripe-go"
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

	//payment
	// fs := http.StripPrefix("/payment/", http.FileServer(http.Dir("./pm/")))
	// r.PathPrefix("/payment/").Handler(fs)
	stripe.Key = "sk_test_51JFVsKKt8F0liU7eQAAKLTZ4mlrjgl53GMDKYXcSxLG7ofODA3V5JfpRqYqKIy2QNnpJRLsoi1WH9Zyi6ON2zemP00cs42GtET"

	// http.Handle("/", http.FileServer(http.Dir(os.Getenv("STATIC_DIR"))))

	// addr := "localhost:4242"
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATH", "DELETE"},
	})

	handler := c.Handler(r)
	// log.Printf("Listening on %s ...", addr)
	// log.Fatal(http.ListenAndServe(addr, nil), handler)

	// http.ListenAndServe(":3000", handlers.CORS()(r))
	http.ListenAndServe(":3000", handler)
}
