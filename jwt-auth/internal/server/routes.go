package server

import (
	"jwt-auth/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *HTTPServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	utils.Error(w, http.StatusNotFound, "The requested endpoint does not exist", nil)
}

func (s *HTTPServer) timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s from %s\n", r.URL.Path, r.Host)
	currentTime := time.Now().Format(time.RFC1123)
	response := map[string]string{"time": currentTime}
	utils.Success(w, http.StatusOK, "Current time fetched successfully", response)
}

func (s *HTTPServer) healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s from %s\n", r.URL.Path, r.Host)
	health := map[string]string{"status": "ok"}
	utils.Success(w, http.StatusOK, "Service is healthy", health)
}

func (s *HTTPServer) Routes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/time", s.timeHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	r.NotFoundHandler = http.HandlerFunc(s.defaultHandler)
	return r
}
