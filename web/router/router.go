package router

import (
	"fmt"
	"net/http"

	"github.com/TestTask/logger"
	"github.com/TestTask/repository"
	"github.com/TestTask/web/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// New return router instance with assigned handlers
func New(log logger.Logger, authRepo repository.AuthRepository) (router *mux.Router,
	headers handlers.CORSOption,
	methods handlers.CORSOption,
	origins handlers.CORSOption,
	err error) {
	router = mux.NewRouter().StrictSlash(true)
	am := middleware.NewAuthMiddleware(log, authRepo)

	router.HandleFunc("/data/get", am.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Here")
	})).Methods(http.MethodGet)

	headers = handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins = handlers.AllowedOrigins([]string{"*"})
	return router, headers, methods, origins, nil
}
