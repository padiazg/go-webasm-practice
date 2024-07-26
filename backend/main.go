package main

import (
	"log"
	"net/http"

	"github.com/padiazg/wasm-app-test/backend/db"
	"github.com/padiazg/wasm-app-test/backend/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db.InitDB()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/users", handlers.GetUsers)
		r.Post("/users", handlers.CreateUser)
		r.Get("/users/{id}", handlers.GetUser)
		r.Put("/users/{id}", handlers.UpdateUser)
		r.Delete("/users/{id}", handlers.DeleteUser)
		r.Post("/login", handlers.Login)
	})

	// Serve index.html for all other routes
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
