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

const (
	SECRET = "#GOLANG"
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

	// Authenication
	router.HandleFunc("/login", JWTAuth(s.userLoginHandler))
	// router.HandleFunc("/logout", JWTAuth(s.userLogoutHandler))

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

	req := AccountSchema{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Error while Unmarshaling!")
		return
	}

	if req.FirstName == "" || req.LastName == "" {
		WriteJSON(w, http.StatusBadRequest, "First & Last name are required!")
		return
	}
	account, err := NewAccount(req.FirstName, req.LastName, req.Password, req.Balance)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Error on Account creation!")
		return
	}

	err = s.Store.CreateAccount(account)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, "Create account failed!")
		return
	}

	tokenString, err := getJWTToken(account)
	if err != nil {
		WriteJSON(w, http.StatusForbidden, "JWT Token err!")
		return
	}

	fmt.Println("TOKEN:", tokenString)

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
		fmt.Println("JWT Authentication...!")
		tokenString := r.Header.Get("x-jwt-token")
		fmt.Println("token:", tokenString)
		_, err := validateJWT(tokenString)
		if err != nil {
			WriteJSON(w, http.StatusForbidden, "Invalid JWT Token!")
			return
		}

		handlerFunc(w, r)
	}
}

func getJWTToken(account *Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":           "JWT-Token",
		"account_number": account.Number,
		"nbf":            time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SECRET))
	return tokenString, err

}

func validateJWT(tokenString string) (*jwt.Token, error) {

	// do validation here
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(SECRET), nil
	})
}

// Handle User Login
func (s *APIServer) userLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	if r.Method != http.MethodPost {
		WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		WriteJSON(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	WriteJSON(w, http.StatusOK, req)
}
