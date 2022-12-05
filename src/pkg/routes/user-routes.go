package routes

import (
	"github.com/DanielT777/Web-Services-Shad/src/pkg/controllers"
	"github.com/gorilla/mux"
)

// UserRoutes ...

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
}
