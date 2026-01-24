package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Port int    `yaml:"port"`
	URL  string `yaml:"url"`
}

type TLS struct {
	UseTLS   bool   `yaml:"use_tls"`
	CertFile string `yaml:"cert_file"`
	KeyFile  string `yaml:"key_file"`
}

type Config struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Storage string `yaml:"storage"`

	Server Server `yaml:"server"`
	TLS    TLS    `yaml:"tls"`
}

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

func PrettyPrintJSON(v any) (err error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		fmt.Println("------------------- Json Start -----------------------")
		fmt.Println(string(b))
		fmt.Println("------------------- Json End -------------------------")
	}
	return err
}

func PrettyPrintYAML(v any) error {
	b, err := yaml.Marshal(v)
	if err == nil {
		fmt.Println("------------------- Yaml Start -----------------------")
		fmt.Println(string(b))
		fmt.Println("------------------- Yaml End -------------------------")
	}
	return err
}

func main() {
	configFile := "config.yaml"
	config, err := LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config:", config)
	fmt.Println("config.Name:", config.Name)
	fmt.Println("config.Type:", config.Type)
	fmt.Println("config.Storage:", config.Storage)
	fmt.Println("config.Server:", config.Server)
	fmt.Println("config.TLS:", config.TLS)

	// PrettyPrint yaml
	err = PrettyPrintYAML(config)
	if err != nil {
		log.Fatal(err)
	}
}
