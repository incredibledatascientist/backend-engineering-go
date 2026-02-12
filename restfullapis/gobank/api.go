package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Addr string `yaml:"addr"`
	// ReadTimeout  time.Duration `yaml:"read_timeout"`
	// WriteTimeout time.Duration `yaml:"write_timeout"`
	// IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

// func NewAPIServer(addr string, handler http.Handler, readTimeout, writeTimeout, idleTimeout time.Duration) *http.Server {
func NewAPIServer(addr string) *APIServer {
	return &APIServer{Addr: addr}
	// return &http.Server{
	// 	Addr: addr,
	// Handler:      handler,
	// ReadTimeout:  readTimeout,
	// WriteTimeout: writeTimeout,
	// IdleTimeout:  idleTimeout,
	// }
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", s.defaultHandler)
	router.HandleFunc("/time", s.timeHandler)
	router.HandleFunc("/health", s.healthHandler)

	// Starting a new server.
	fmt.Println("Server is listening on addr:", s.Addr)
	err := http.ListenAndServe(s.Addr, router)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *APIServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Thanks for visiting...")
}

func (s *APIServer) timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format(time.RFC1123)
	body := fmt.Sprintf("Current time is: %v\n", t)
	fmt.Fprintf(w, "%s", body)
}

func (s *APIServer) healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ok"}`)
}
