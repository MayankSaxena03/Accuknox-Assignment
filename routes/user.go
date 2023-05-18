package routes

import (
	"github.com/MayamkSaxena03/AccuknoxAssignment/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
}
