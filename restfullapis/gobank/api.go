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
	return &APIServer{
		Addr:         cfg.Addr,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", s.defaultHandler)
	router.HandleFunc("/time", s.timeHandler)
	router.HandleFunc("/health", s.healthHandler)
	router.HandleFunc("/account/{id}", s.getAccountHandler)
	// router.HandleFunc("/view", s.getAccountHandler)
	router.HandleFunc("/add", s.addAccountHandler)

	// Starting a new server.
	fmt.Println("Server is listening on addr:", s.Addr)
	// err := http.ListenAndServe(s.Addr, router) // For practice only

	server := &http.Server{
		Addr:         s.Addr,
		Handler:      router,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		IdleTimeout:  s.IdleTimeout,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
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

func (s *APIServer) addAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Method not allowed!|")
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
		http.Error(w, "Error while Unmarshal", http.StatusBadRequest)
		return
	}

	log.Println("account", account)
	if account.FirstName != "" && account.LastName != "" {
		newAcc := NewAccount(account.FirstName, account.LastName, account.Balance)
		CUSTOMERS = append(CUSTOMERS, *newAcc)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
	} else {
		http.Error(w, "First & Last name are required. ", http.StatusBadRequest)
		return
	}
}

func (s *APIServer) getAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("vars:", vars)

	log.Println("Serving:", r.URL.Path, "from", r.Host)
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Method not allowed!|")
		return
	}

	v, ok := vars["id"]
	var account Account
	if ok {
		fmt.Printf("v:%v, %T", v, v)

		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, "Invalid id")
			return
		}
		for _, acc := range CUSTOMERS {
			if acc.Id == id {
				account = acc
			}
		}
	}

	if account.Id == 0 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Account doesn't exists!")
		return
	}

	accountJson, err := json.Marshal(account)
	if err != nil {
		http.Error(w, "Error while marshal", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v\n", string(accountJson))

	// d, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Error on read", http.StatusBadRequest)
	// 	return
	// }

	// var account Account
	// err = json.Unmarshal(d, &account)

	// if err != nil {
	// 	http.Error(w, "Error while marshal", http.StatusBadRequest)
	// 	return
	// }

	// log.Println("account", account)
	// if account.FirstName != "" && account.LastName != "" {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Header().Set("Content-Type", "application/json")
	// } else {
	// 	http.Error(w, "First & Last name are required. ", http.StatusBadRequest)
	// 	return
	// }
}
