package main

import (
	f "fmt"
	// "project_web/structs"
	"net/http"
	"log"
)


func handler(w http.ResponseWriter, r *http.Request) {

	// f.Println(r.URL.Query())
	queryParameters := r.URL.Query()
	for key, value := range queryParameters {
		f.Printf("Key: %s val: %s\n", key, value)
	}

	f.Fprint(w, "Hi this is a response", r.URL.Path[1:])
}



func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
