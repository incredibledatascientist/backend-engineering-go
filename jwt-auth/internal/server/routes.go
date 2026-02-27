package server

import (
	"jwt-auth/internal/middleware"
	"jwt-auth/internal/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *HTTPServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	utils.Error(w, http.StatusNotFound, "The requested endpoint does not exist", nil)
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

	r.HandleFunc("/time", s.timeHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.NotFoundHandler = http.HandlerFunc(s.defaultHandler)

	// User Routes
	// r.HandleFunc("/users/signup", h.SignupHandler).Methods(http.MethodPost)
	// r.HandleFunc("/users/login", h.LoginHandler).Methods(http.MethodPost)
	// r.HandleFunc("/users", h.GetUsersHandler).Methods(http.MethodGet)
	// r.HandleFunc("/users/{user_id}", h.GetUserHandler).Methods(http.MethodGet)

	return r
}
