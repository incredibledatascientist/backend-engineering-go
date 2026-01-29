package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// type APIServer struct {
// 	Addr         string
// 	ReadTimeout  time.Duration
// 	WriteTimeout time.Duration
// }

// func NewAPIServer(port int, read_timeout, write_timeout time.Duration) {
// 	return &APIServer{
// 		Addr:         fmt.Sprintf(":", port),
// 		ReadTimeout:  read_timeout,
// 		WriteTimeout: write_timeout,
// 	}
// }

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusNotFound)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"status":"ok"}`))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format(time.RFC1123)
	body := fmt.Sprintf("Current time is: %v\n", t)
	fmt.Fprintf(w, "%s", body)
}

func main() {
	mux := http.NewServeMux()

	// Handle Endpoints
	mux.Handle("/", http.HandlerFunc(defaultHandler))
	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/health", http.HandlerFunc(healthHandler))

	server := http.Server{
		// Addr:         ":8080", // Every time asking allow
		// Addr:         "0.0.0.0:8080",

		// Addr:         "127.0.0.1:8080", // Not aski for allow network.
		Addr:         "localhost:8080", // Not aski for allow network.
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
		Handler:      mux,
	}

	// fmt.Println("Server is running...")
	fmt.Println("Ready to serve at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
	// server.ListenAndServeTLS() // Enable HTTPS

	// http.ListenAndServe()
	// 	server := &http.Server{
	//     Addr:    ":8080",
	//     Handler: nil, // means DefaultServeMux
	// }
	// server.ListenAndServe()

}
