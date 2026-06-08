package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mayankdev-oss/go-bookstore-sql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
