package main

import (
	"bankproject/database"
	"bankproject/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	DB := database.Init()
	router := mux.NewRouter()
	routers.SetupRoutes(DB, router)
	log.Println("API is running!")
	http.ListenAndServe(":3000", router)
}
