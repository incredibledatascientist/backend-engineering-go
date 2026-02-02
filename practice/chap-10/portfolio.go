package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name   string `yaml:"name"`
	Server Server `yaml:"server"`
	// Storage Storage `yaml:"storage"`
	// TLS    TLS    `yaml:"tls"`
}

type Server struct {
	Addr         string        `yaml:"addr"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`
	IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

// type Storage struct {
// 	Name     string `yaml:"name"`
// 	host     string `yaml:"host"`
// 	Port     int    `yaml:"port"`
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// }

type Contacts struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

var contacts []Config

func LoadConfig(configFile string) (Config, error) {
	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("[LoadConfig] (%v)", err)
	}
	var config Config
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("[LoadConfig] (%v)", err)
	}
	return config, nil
}

// func PrettyPrintJSON(v any) (err error) {
// 	b, err := json.MarshalIndent(v, "", "\t")
// 	if err == nil {
// 		fmt.Println("------------------- Json Start -----------------------")
// 		fmt.Println(string(b))
// 		fmt.Println("------------------- Json End -------------------------")
// 	}
// 	return err
// }

// Get the configs from yaml

func NewAPIServer(addr string, handler http.Handler, readTimeout, writeTimeout, idleTimeout time.Duration) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}
}

func PrettyPrintYAML(v any) {
	b, err := yaml.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("------------------- Yaml Start -----------------------")
	fmt.Println(string(b))
	fmt.Println("------------------- Yaml End -------------------------")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Thanks for visiting...")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status":"ok"}`)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format(time.RFC1123)
	body := fmt.Sprintf("Current time is: %v\n", t)
	fmt.Fprintf(w, "%s", body)
}

func addContactsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	// Method check
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Content-Type check
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "content-type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	defer r.Body.Close()

	// Decode JSON
	var contact Contacts
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Validate fields
	if strings.TrimSpace(contact.Name) == "" ||
		strings.TrimSpace(contact.Email) == "" ||
		strings.TrimSpace(contact.Message) == "" {

		http.Error(w, "name, email and message are required", http.StatusBadRequest)
		return
	}

	// Business logic
	log.Println("New contact:", contact)

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}

func main() {
	configFile := "config.yaml"
	cfg, err := LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Config server:", cfg.Server)
	// fmt.Println("Config name:", cfg.Name)
	// // PrettyPrint yaml
	// PrettyPrintYAML(cfg)

	mux := http.NewServeMux()

	mux.Handle("/health", http.HandlerFunc(healthHandler))
	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	// Contact API
	mux.Handle("/contacts/add", http.HandlerFunc(addContactsHandler))

	// // Starting a new server.
	server := NewAPIServer(
		cfg.Server.Addr, mux,
		cfg.Server.ReadTimeout,
		cfg.Server.WriteTimeout,
		cfg.Server.IdleTimeout,
	)
	fmt.Println("Server is listening on addr:", cfg.Server.Addr)
	server.ListenAndServe()
}
