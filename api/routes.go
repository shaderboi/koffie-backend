package api

import (
	"github.com/gorilla/mux"
	"github.com/shaderboi/koffie-backend/api/controllers"
	"github.com/shaderboi/koffie-backend/api/middleware"
	"log"
	"net/http"
)

func Routes() {
	r := mux.NewRouter()
	r.Use(middleware.AuthMiddleware)
	r.HandleFunc("/api/v1/products", controllers.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/v1/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/api/v1/products/{code}", controllers.GetProduct).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
