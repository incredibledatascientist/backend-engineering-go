package utils

import (
	"encoding/json"
	"fmt"
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
