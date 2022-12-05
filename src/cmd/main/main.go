package main

import (
	"log"
	"net/http"

	"github.com/DanielT777/Web-Services-Shad/src/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterUserRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
