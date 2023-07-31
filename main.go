package main

import (
	"github.com/gorilla/mux"
	"github.com/nikitwei/alfa-shop-api/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterProductsRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
