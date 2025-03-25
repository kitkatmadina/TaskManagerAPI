package main

import (
	"log"
	"net/http"

	"github.com/kitkatmadina/TaskManagerAPI/db"
	"github.com/kitkatmadina/TaskManagerAPI/routes"
)

func main() {
	db.Init()
	r := routes.SetRoutes()

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
