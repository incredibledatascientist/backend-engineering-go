package main

import (
	"flag"
	"fmt"
)

func main() {
	// port := flag.Int("port", 8080, "Port number")
	// useTLS := flag.Bool("tls", false, "Use TLS")
	// configFile := flag.String("config", "config.yaml", "Configuration file.")
	// flag.Parse()

	// fmt.Println("Port:", *port)
	// fmt.Println("UserTLS:", *useTLS)
	// fmt.Println("ConfigFile:", *configFile)

	// Optimized way with var, No need to de-refference
	var port int
	var useTLS bool
	var configFile string
	flag.IntVar(&port, "port", 8080, "Port number")
	flag.BoolVar(&useTLS, "tls", false, "Use TLS")
	flag.StringVar(&configFile, "config", "config.yaml", "Configuration file.")
	flag.Parse()

	fmt.Println("Port:", port)
	fmt.Println("UserTLS:", useTLS)
	fmt.Println("ConfigFile:", configFile)

	// Usage : go run .\flags.go --help
	// 	-config string
	//         Configuration file. (default "config.yaml")
	//   -port int
	//         Port number (default 8080)
	//   -tls
	//         Use TLS
}
