package main

// Default GoLang Imports
// Dependencies Imports
// Handler Functions Imports
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/vincent-pantaleon/orders-api/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/get", handlers.GetHandler)
	router.Get("/post", handlers.PostHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error Running Server with Port: 3000", err)
	}
}
