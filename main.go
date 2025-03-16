package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
    w.Header().Set("Content-Type", "text/html")
    tpl, err := template.ParseFiles(filepath)
    if err != nil {
        log.Printf("Error parsing template: %v", err)
        http.Error(w, "there was an error parsing the template", http.StatusInternalServerError)
        return
    }
    err = tpl.Execute(w, nil)
    if err != nil {
        log.Printf("Error executing template: %v", err)
        http.Error(w, "there was an error executing the template", http.StatusInternalServerError)
        return
    }
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    tplPath := filepath.Join("templates", "home.gohtml")
    executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    // w.Write([]byte("Contact us at support@photodump.com"))
    tplPath := filepath.Join("templates", "contact.gohtml")
    executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Frequently Asked Questions"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    userID := chi.URLParam(r, "userID")
    w.Write([]byte(fmt.Sprintf("User ID: %s", userID)))
}

func main() {
    r := chi.NewRouter()
    
    // Add some middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Define routes
    r.Get("/", handlerFunc)
    r.Get("/contact", contactHandler)
    r.Get("/faq", faqHandler)
    r.Get("/user/{userID}", userHandler) // Route with URL parameter

    // Custom 404 handler
    r.NotFound(func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Page not found", http.StatusNotFound)
    })

    // Start the server
    fmt.Println("Server is running on port 3000")
    http.ListenAndServe(":3000", r)
}