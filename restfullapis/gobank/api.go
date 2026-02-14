package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func NewAPIServer(cfg ServerConfig) *APIServer {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	return &APIServer{
		Server: cfg,
		Addr:   addr,
	}
}

func (s *APIServer) Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", s.defaultHandler)
	router.HandleFunc("/time", s.timeHandler)
	router.HandleFunc("/health", s.healthHandler)
	router.HandleFunc("/account/{id}", s.getAccountHandler)
	// router.HandleFunc("/view", s.getAccountHandler)
	router.HandleFunc("/add", s.addAccountHandler)

	return router
}

func (s *APIServer) Run() error {
	// Starting a new server.
	fmt.Println("Server is listening on addr:", s.Addr)
	server := &http.Server{
		Addr:         s.Addr,
		Handler:      s.Routes(),
		ReadTimeout:  s.Server.ReadTimeout,
		WriteTimeout: s.Server.WriteTimeout,
		IdleTimeout:  s.Server.IdleTimeout,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("failed to encode response")
		return
	}
}

func (s *APIServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	// w.WriteHeader(http.StatusNotFound)
	// w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, "Thanks for visiting...")
	WriteJSON(w, http.StatusNotFound, "Thanks for visiting...")
}

func (s *APIServer) timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format(time.RFC1123)
	body := fmt.Sprintf("Current time is: %v\n", t)
	WriteJSON(w, http.StatusOK, body)
}

func (s *APIServer) healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	health := map[string]string{"status": "ok"}
	WriteJSON(w, http.StatusOK, health)
}

func (s *APIServer) addAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	if r.Method != http.MethodPost {
		WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error on read", http.StatusBadRequest)
		return
	}

	var account Account
	err = json.Unmarshal(d, &account)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Error while Unmarshaling!")
		return
	}

	log.Println("account", account)
	if account.FirstName != "" && account.LastName != "" {
		newAcc := NewAccount(account.FirstName, account.LastName, account.Balance)
		CUSTOMERS = append(CUSTOMERS, *newAcc)
		WriteJSON(w, http.StatusCreated, "Account created successfully!")
	} else {
		WriteJSON(w, http.StatusBadRequest, "First & Last name are required!")
		return
	}
}

func (s *APIServer) getAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	vars := mux.Vars(r)

	if r.Method != http.MethodGet {
		WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	v, ok := vars["id"]
	var account Account
	if ok {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, "Invalid id!")
			return
		}
		for _, acc := range CUSTOMERS {
			if acc.Id == id {
				account = acc
			}
		}
	}

	if account.Id == 0 {
		WriteJSON(w, http.StatusNotFound, "Account doesn't exists!")
		return
	}

	WriteJSON(w, http.StatusOK, account)
}
