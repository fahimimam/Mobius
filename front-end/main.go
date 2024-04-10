package main

import (
	"fmt"
	"frontend/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Serve static files
	directoryPath := "./static"
	// Check if the directory exists
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		return
	}

	// Create a file server handler to serve the directory's contents
	fileServer := http.FileServer(http.Dir(directoryPath))
	// Create a new HTTP server and handle requests
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// handle other routes
	router.Get("/home", landingPage)
	router.Get("/register", register)

	srv := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	server.HandleServer(srv)
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	render(w, "test.page.gohtml")
}

func register(w http.ResponseWriter, r *http.Request) {
	partials := []string{
		"./templates/base.layout.gohtml",
		"./templates/footer.partial.gohtml",
	}
	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./templates/%s", "register.page.gohtml"))

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
		"./templates/base.layout.gohtml",
		"./templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./templates/%s", t))

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
