package routes

import (
	"github.com/MayamkSaxena03/Accuknox/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
}
