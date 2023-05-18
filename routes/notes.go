package routes

import (
	"github.com/MayamkSaxena03/Accuknox/controllers"
	"github.com/gorilla/mux"
)

func NotesRoutes(router *mux.Router) {
	router.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
	router.HandleFunc("/notes", controllers.DeleteNote).Methods("DELETE")
}
