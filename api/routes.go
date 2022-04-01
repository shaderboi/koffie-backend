package api

import (
	"github.com/gorilla/mux"
	"github.com/shaderboi/koffie-backend/api/controllers"
	"log"
	"net/http"
)

func Routes() {
	r := mux.NewRouter()
	r.HandleFunc("/api/products", controllers.GetAllProducts).Methods("GET")
	r.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
