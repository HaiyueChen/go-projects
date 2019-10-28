package controllers

import (
	f "fmt"
	"html/template"
	"net/http"
	"strings"
)

func parseTemplate(path string) (*template.Template, error) {
	t, err := template.ParseFiles("html/index.html")

	if err != nil {
		return nil, err

	}
	return t, nil
	// t.Execute(w, nil)
}

//RootController for the server root
func RootController(w http.ResponseWriter, r *http.Request) {

	// queryParameters := r.URL.Query()
	// for key, value := range queryParameters {
	// 	f.Printf("Key: %s val: %s\n", key, value)
	// }
	controllers := make(map[string]func(http.ResponseWriter, *http.Request){
		"edit": editController})

	controllers["edit"]()

	urlPath := r.URL.Path
	if urlPath == "/" || urlPath == "/index" || urlPath == "index.html" {
		template, err := parseTemplate("html/index.html")
		if err != nil {

		}
		template.Execute(w, nil)
		return
	}

	urlPath = urlPath[1:]
	controllerKey := strings.Split(urlPath, "/")[0]
	nextController := controllers[controllerKey]
	go nextController(w, r)
	return
}

func editController(w http.ResponseWriter, r *http.Request) {
	restPath := r.URL.Path[len("/edit/"):]
	f.Println(restPath)
	if restPath == "" {

	} else {
		
	}

}
