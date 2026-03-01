package server

import (
	"bookstore/internal/handler"
	"bookstore/internal/middleware"
	"bookstore/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(router *mux.Router, h *handler.BookHandler) {
	router.HandleFunc("/books", h.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books", h.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)
}

func (s *HTTPServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	err := fmt.Sprintf("The requested endpoint (%v) does not exist", r.RequestURI)
	utils.Error(w, http.StatusNotFound, err, nil)
}

func (s *HTTPServer) timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	response := map[string]string{"time": currentTime}
	utils.Success(w, http.StatusOK, "Current time fetched successfully", response)
}

func (s *HTTPServer) healthHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]string{"status": "ok"}
	utils.Success(w, http.StatusOK, "Service is healthy", health)
}

func (s *HTTPServer) Routes() http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.Logging)
	bookHandler := handler.NewBookHandler(s.store)

	// api := r.PathPrefix("/api/v1").Subrouter()
	RegisterBookStoreRoutes(r, bookHandler)

	r.HandleFunc("/time", s.timeHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.NotFoundHandler = http.HandlerFunc(s.defaultHandler)

	return r
}
