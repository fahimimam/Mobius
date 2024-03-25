package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	staticFileDirectory := http.Dir("/assets/css/")
	log.Println("Got Directory as ", staticFileDirectory)
	staticFileHandler := http.FileServer(staticFileDirectory) //http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	http.Handle("/assets/css/", staticFileHandler)

	http.HandleFunc("/home", landingPage)
	http.HandleFunc("/register", register)

	fmt.Println("Starting front end service on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Panic(err)
	}
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	render(w, "test.page.gohtml")
}

func register(w http.ResponseWriter, r *http.Request) {
	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}
	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", "register.page.gohtml"))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Println("Executed Templates")
}

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
