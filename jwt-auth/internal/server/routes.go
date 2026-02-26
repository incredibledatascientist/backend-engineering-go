package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("failed to encode response")
		return
	}
}

func (s *HTTPServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	WriteJSON(w, http.StatusNotFound, "Thanks for visiting...")
}

func (s *HTTPServer) timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format(time.RFC1123)
	body := fmt.Sprintf("Current time is: %v\n", t)
	WriteJSON(w, http.StatusOK, body)
}

func (s *HTTPServer) healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	health := map[string]string{"status": "ok"}
	WriteJSON(w, http.StatusOK, health)
}

func (s *HTTPServer) Routes() http.Handler {
	router := mux.NewRouter()

	// router.HandleFunc("/", s.defaultHandler)
	router.HandleFunc("/time", s.timeHandler)
	router.HandleFunc("/health", s.healthHandler)

	return router
}
