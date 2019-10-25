package handlers

import (
	f "fmt"
	"net/http"
)

//RootHandler for the server root
func RootHandler(w http.ResponseWriter, r *http.Request) {

	// f.Println(r.URL.Query())
	queryParameters := r.URL.Query()
	for key, value := range queryParameters {
		f.Printf("Key: %s val: %s\n", key, value)
	}

	f.Fprint(w, "Hi this is a response", r.URL.Path[1:])
}
