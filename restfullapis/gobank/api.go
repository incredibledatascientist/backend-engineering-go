package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func NewAPIServer(cfg ServerConfig, store Storage) *APIServer {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	return &APIServer{
		Server: cfg,
		Addr:   addr,
		Store:  store,
	}
}

func (s *APIServer) Routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", s.defaultHandler)
	router.HandleFunc("/time", s.timeHandler)
	router.HandleFunc("/health", s.healthHandler)
	router.HandleFunc("/account", s.getAccountHandler)
	router.HandleFunc("/account/{id}", s.getAccountHandler)
	router.HandleFunc("/account/delete/{id}", s.deleteAccountHandler)
	// router.HandleFunc("/view", s.getAccountHandler)
	router.HandleFunc("/add", s.addAccountHandler)

	router.HandleFunc("/transfer", JWTAuth(s.transferAccountHandler))

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

	// var accSchema AccountSchema
	accSchema := AccountSchema{}
	err := json.NewDecoder(r.Body).Decode(&accSchema)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Error while Unmarshaling!")
		return
	}

	if accSchema.FirstName == "" || accSchema.LastName == "" {
		WriteJSON(w, http.StatusBadRequest, "First & Last name are required!")
		return
	}
	account := NewAccount(accSchema.FirstName, accSchema.LastName, accSchema.Balance)

	err = s.Store.CreateAccount(account)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Create account failed!")
		return
	}

	// Account created successfully!
	WriteJSON(w, http.StatusCreated, account)

}

func (s *APIServer) getAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	vars := mux.Vars(r)

	if r.Method != http.MethodGet {
		WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	v, ok := vars["id"]
	if !ok || v == "" {
		accounts, err := s.Store.GetAllAccount()
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, "Error on db query!")
			return
		}
		WriteJSON(w, http.StatusOK, accounts)
		return
	}

	id, err := getID(r)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	account, err := s.Store.GetAccount(id)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, "Account doesn't exist!")
		return
	}

	WriteJSON(w, http.StatusOK, account)
}
func (s *APIServer) deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	if r.Method != http.MethodDelete {
		WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	id, err := getID(r)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.Store.DeleteAccount(id)
	if err != nil {
		WriteJSON(w, http.StatusNotFound, "Account not found")
		return
	}

	// WriteJSON(w, http.StatusNoContent, "") // Delete successfull but no message
	WriteJSON(w, http.StatusOK, "Account deleted successfully!")
}

func getID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	v, ok := vars["id"]
	if !ok || v == "" {
		return 0, fmt.Errorf("Id not present")
	}

	id, err := strconv.Atoi(v)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("Invalid id %s", v)
	}
	return id, nil
}

func (s *APIServer) transferAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	if r.Method != http.MethodPost {
		WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	// var transferSchema
	transferSchema := TransferSchema{}
	err := json.NewDecoder(r.Body).Decode(&transferSchema)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Error while Unmarshaling!")
		return
	}

	if transferSchema.ToAccount == "" || transferSchema.Amount == 0 {
		WriteJSON(w, http.StatusBadRequest, "Account number and amount are required!")
		return
	}

	WriteJSON(w, http.StatusOK, transferSchema)
}

// ---------------------- JWT Authentication ----------------------
func JWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
		fmt.Println("JWT Authentication...!")
	}
}

func validateJWT() (*jwt.Token, error) {

}
