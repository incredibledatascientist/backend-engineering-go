package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	Addr string `yaml:"addr"`
	// ReadTimeout  time.Duration `yaml:"read_timeout"`
	// WriteTimeout time.Duration `yaml:"write_timeout"`
	// IdleTimeout  time.Duration `yaml:"idle_timeout"`
}

// type Storage struct {
// 	Name     string `yaml:"name"`
// 	host     string `yaml:"host"`
// 	Port     int    `yaml:"port"`
// 	Username string `yaml:"username"`
// 	Password string `yaml:"password"`
// }

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
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	// // Starting a new server.
	server := NewAPIServer(cfg.Server.Addr, mux, 10*time.Second, 10*time.Second, 30*time.Second)
	fmt.Println("Server is listening on addr:", cfg.Server.Addr)
	server.ListenAndServe()
}
