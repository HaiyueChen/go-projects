package main

import (
	f "fmt"
	// "project_web/structs"
	"log"
	"net/http"
	"project_web/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	port := ":3000"
	f.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
