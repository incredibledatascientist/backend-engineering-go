package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// JSON/YAML ONLY MARSHAL/UNMARSHAL EXPORTED FIELDS [ALWAYS USE Capital Letter for FIELDS]
type Server struct {
	// port int // Not exported
	// addr string // Not exported

	Port int    `yaml:"port"`
	Addr string `yaml:"addr"`
}

type Config struct {
	// Server        // Direct embedding will flat the struct it won't work under server.
	Server Server `yaml:"server"`
	Name   string `yaml:"name"`
	Type   string `yaml:"type"`
}

func main() {
	// --------------- Marshal YAML File and Read ----------------
	fmt.Println("--------------- Marshal YAML File and Read ----------------")

	// Read Config file
	configFile := "config.yaml"
	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("File bytes:", file)
	fmt.Println("Raw File string:", string(file))

	var server Config
	err = yaml.Unmarshal(file, &server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Struct File:", server)

	// --------------- Marshal YAML File and Write ----------------
	fmt.Println("--------------- Marshal YAML File and Write ----------------")
	// Write into yaml file
	serverBytes, err := yaml.Marshal(&server)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("serverBytes:", serverBytes)
	fmt.Println("serverString:", string(serverBytes))

	// Write into yaml file
	err = os.WriteFile(configFile, serverBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config written into yaml file.")
	// writeFile, err := os.Create("config.yaml")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer writeFile.Close()
	// writeFile.Write(ymlBytes)

}
