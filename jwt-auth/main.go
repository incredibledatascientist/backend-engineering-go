package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"jwt-auth/internal/config"
	"log"

	"github.com/goccy/go-yaml"
)

func PrettyPrintJSON(v any) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		fmt.Println("------------------- Json Start -----------------------")
		fmt.Print(string(b))
		fmt.Println("------------------- Json End -------------------------")
	}
	return err
}

func PrettyPrintYAML(v any) {
	b, err := yaml.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("------------------- Yaml Start -----------------------")
	fmt.Print(string(b))
	fmt.Println("------------------- Yaml End -------------------------")
}

// func NewAPIServer(cfg models.ServerConfig, store Storage) *models.APIServer {
func NewHTTPServer(cfg config.Config) *config.HTTPServer {
	return &config.HTTPServer{
		Addr:         cfg.Server.Addr,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}
}

// func (s *APIServer) Run() error {
// 	// Starting a new server.
// 	fmt.Println("Server is listening on addr:", s.Addr)
// 	server := &http.Server{
// 		Addr:         s.Addr,
// 		Handler:      s.Routes(),
// 		ReadTimeout:  s.Server.ReadTimeout,
// 		WriteTimeout: s.Server.WriteTimeout,
// 		IdleTimeout:  s.Server.IdleTimeout,
// 	}
// 	err := server.ListenAndServe()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func main() {
	configFile := flag.String("config", "configs/local.yaml", "Configuration file.")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	PrettyPrintYAML(cfg)

	// store, err := NewPostgresStore()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Create table
	// err = store.createAccountTable()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server := NewAPIServer(cfg, store)
	// server := NewAPIServer(cfg)
	// if err := server.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
