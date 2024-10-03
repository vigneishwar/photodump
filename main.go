package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to the homepage!"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Contact us at support@photodump.com"))
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