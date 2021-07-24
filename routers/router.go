package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quang2906/book_store_be/controller"
	"github.com/quang2906/book_store_be/middleware"
	"github.com/quang2906/book_store_be/util"
)

// func ConfigCategoryRouter(router *mux.Router) {
// 	category := router.PathPrefix("/v1/categories").Subrouter()
// 	category.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllCategory)
// 	category.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetCategoryById)
// 	category.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateCategory)
// 	category.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateCategory)
// 	category.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteCategoryById)
// }

// func ConfigProductsRouter(router *mux.Router) {
// 	product := router.PathPrefix("/v1/products").Subrouter()

// 	product.Path("").Methods(http.MethodGet).HandlerFunc(controller.SearchProduct)
// 	product.Path("/sort").Methods(http.MethodGet).HandlerFunc(controller.SortProduct)
// 	// product.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllProducts)
// 	product.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetProductById)
// 	product.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateProduct)
// 	product.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateProductById)
// 	product.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteProductById)

// }

func ConfigUserRouter(router *mux.Router) {
	user := router.PathPrefix("/v1/users").Subrouter()
	// 	user.Path("/login").Methods(http.MethodPost).HandlerFunc(controller.SignIn)
	// 	user.Path("/logout").Methods(http.MethodPost).HandlerFunc(controller.Logout)

	// 	user.Path("/admin").Methods(http.MethodGet).HandlerFunc(middleware.IsAuthorized(controller.GetAllUsers))
	// 	user.Path("/admin/{id}").Methods(http.MethodGet).HandlerFunc(middleware.IsAuthorized(controller.GetUserById))
	user.Path("").Methods(http.MethodPost).HandlerFunc(controller.CreateUser)
	// 	user.Path("/admin/{id}").Methods(http.MethodPut).HandlerFunc(middleware.IsAuthorized(controller.UpdateUser))
	// 	user.Path("/admin/{id}").Methods(http.MethodDelete).HandlerFunc(middleware.IsAuthorized(controller.DeleteUserById))

}

// func ConfigOrderRouter(router *mux.Router) {
// 	order := router.PathPrefix("/v1/orders").Subrouter()
// 	order.Path("").Methods(http.MethodGet).HandlerFunc(controller.GetAllOrder)
// 	order.Path("/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetOrderById)
// 	order.Path("/user").Methods(http.MethodPost).HandlerFunc(middleware.IsAuthorized(controller.))
// 	order.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controller.UpdateOrderById)
// 	order.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controller.DeleteOrderById)
// }

func ConfigFileRouter(router *mux.Router) {
	r := router.PathPrefix("/v1").Subrouter()
	r.Path("/uploads").Methods(http.MethodPost).HandlerFunc(util.UploadFile)
	images := http.StripPrefix("/images/", http.FileServer(http.Dir("./public/")))
	r.PathPrefix("/images/").Handler(images)
}
func ConfigRouter(router *mux.Router) {
	customerURL := router.PathPrefix("/api/v1").Subrouter()
	customerURL.Path("/search").Methods(http.MethodGet).HandlerFunc(controller.SearchProduct)
	customerURL.Path("/sort").Methods(http.MethodGet).HandlerFunc(controller.SortProduct)
	customerURL.Path("/categories").Methods(http.MethodGet).HandlerFunc(controller.GetAllCategory)
	customerURL.Path("/category/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetCategoryById)
	customerURL.Path("/produt/{id}").Methods(http.MethodGet).HandlerFunc(controller.GetProductById)

	adminURL := router.PathPrefix("/admin").Subrouter()

	router.HandleFunc("/create-payment-intent", controller.HandleCreatePaymentIntent).Methods(http.MethodPost)

	router.Path("/login").Methods(http.MethodPost).HandlerFunc(controller.SignIn)
	router.Path("/logout").Methods(http.MethodPost).HandlerFunc(controller.Logout)

	adminURL.Path("/users").Methods(http.MethodGet).HandlerFunc(middleware.IsAuthorized(controller.GetAllUsers))
	adminURL.Path("/user/{id}").Methods(http.MethodGet).HandlerFunc(middleware.IsAuthorized(controller.GetUserById))
	adminURL.Path("/user").Methods(http.MethodPost).HandlerFunc(middleware.IsAuthorized(controller.CreateUser))
	adminURL.Path("/user/{id}").Methods(http.MethodPut).HandlerFunc(middleware.IsAuthorized(controller.UpdateUser))
	adminURL.Path("/user/{id}").Methods(http.MethodDelete).HandlerFunc(middleware.IsAuthorized(controller.DeleteUserById))

	adminURL.Path("/product").Methods(http.MethodPost).HandlerFunc(middleware.IsAuthorized(controller.CreateProduct))
	adminURL.Path("/product/{id}").Methods(http.MethodPut).HandlerFunc(middleware.IsAuthorized(controller.UpdateProductById))
	adminURL.Path("/product/{id}").Methods(http.MethodDelete).HandlerFunc(middleware.IsAuthorized(controller.DeleteProductById))

	adminURL.Path("/category").Methods(http.MethodPost).HandlerFunc(middleware.IsAuthorized(controller.CreateCategory))
	adminURL.Path("/category/{id}").Methods(http.MethodPut).HandlerFunc(middleware.IsAuthorized(controller.UpdateCategory))
	adminURL.Path("/category/{id}").Methods(http.MethodDelete).HandlerFunc(middleware.IsAuthorized(controller.DeleteCategoryById))

	adminURL.Path("/order/{id}").Methods(http.MethodPut).HandlerFunc(middleware.IsAuthorized(controller.UpdateOrderById))
	adminURL.Path("/order/{id}").Methods(http.MethodDelete).HandlerFunc(middleware.IsAuthorized(controller.DeleteOrderById))

	adminURL.Path("/orders").Methods(http.MethodGet).HandlerFunc(middleware.IsAuthorized(controller.GetAllOrder))
	adminURL.Path("/order/{id}").Methods(http.MethodGet).HandlerFunc(middleware.IsAuthorized(controller.GetOrderById))

	userURL := router.PathPrefix("/user").Subrouter()

	userURL.Path("").Methods(http.MethodPost).HandlerFunc(middleware.IsAuthorized(controller.CreateOrder))

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:*", "http://192.168.8.235:*", "http://118.69.210.244:*"},
	// 	AllowCredentials: true,
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "PATH", "DELETE"},
	// })
	// handler := c.Handler(router)
}
