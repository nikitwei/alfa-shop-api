package routes

import (
	"github.com/gorilla/mux"
	"github.com/nikitwei/alfa-shop-api/pkg/controllers"
)

var RegisterProductsRoutes = func(router *mux.Router) {
	router.HandleFunc("/product", controllers.AddProduct).Methods("POST")
	router.HandleFunc("/product", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.GetProductID).Methods("GET")
	router.HandleFunc("/product/{productId}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{productId}", controllers.DeleteProduct).Methods("DELETE")
}
