package routes

import (
	"github.com/gorilla/mux"
	"github.com/starlingvibes/Simple-CRUD-MonGo/controllers"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/user/{userID}", controllers.GetAUser()).Methods("GET")
	router.HandleFunc("/user/{userID}", controllers.EditAUser()).Methods("PUT")
	router.HandleFunc("/user/{userID}", controllers.DeleteAUser()).Methods("DELETE")
	router.HandleFunc("/users", controllers.GetAllUsers()).Methods("GET")
}