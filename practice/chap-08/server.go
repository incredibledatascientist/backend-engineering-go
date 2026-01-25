package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Greeting...")
	fmt.Fprintf(w, "Welcome to golang server ...")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello!")
	w.Write([]byte("Hello, production Go 👋"))
	fmt.Println("Hello!")

	fmt.Println("---------------")
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println(r.Body)
	fmt.Println("---------------")
}

func main() {

	// // Default server Endpoints
	// http.HandleFunc("/greet", greeting)
	// http.HandleFunc("/hello", handler)
	// err := http.ListenAndServe(":8080", nil)
	// Default router automaticall add all the routes ex. debug/pprof

	// Custom servermux for production
	mux := http.NewServeMux()

	// Custom Server End points
	mux.HandleFunc("/greet", greeting)
	mux.HandleFunc("/hello", handler)

	// pprof routes (explicit)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	fmt.Println("Server is running...")
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
