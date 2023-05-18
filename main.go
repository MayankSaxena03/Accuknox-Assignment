package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MayamkSaxena03/Accuknox/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := mux.NewRouter()

	routes.UserRoutes(router)
	routes.NotesRoutes(router)

	fmt.Println("Server is running on port " + port)
	http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, router))
}
