package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type config struct {
	ApiKey string `json:"apikey"`
}

type MainBody struct {
	Temp float64 `json:"temp"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type System struct {
	Country string `json:"country"`
}

type WeatherData struct {
	Name  string      `json:"name"`
	Main  MainBody    `json:"main"`
	Coord Coordinates `json:"coord"`
	Sys   System      `json:"sys"`
}

func loadConfig(filepath string) (config, error) {
	cfg := config{}

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func main() {
	// cfg, err := loadConfig("config.json")
	// if err != nil {
	// 	log.Fatal("[main] err:", err.Error())
	// }

	// fmt.Println("cfg:", cfg.ApiKey)

	// Handlers
	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/weather/{city}", weatherHandler).Methods(http.MethodGet)

	server := http.Server{
		Addr:    "localhost:8888",
		Handler: r,
	}

	fmt.Println("Server is running on addr:", server.Addr)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("[server] err:", err.Error())
	}

}

func query(city string) (WeatherData, error) {
	fmt.Println("-----------------------------")
	cfg, err := loadConfig("config.json")
	if err != nil {
		log.Fatal("[query] err:", err.Error())
	}

	fmt.Println("apiKey:", cfg.ApiKey)

	apiURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, cfg.ApiKey)
	// d := WeatherData{
	// 	City: city,
	// }

	client := &http.Client{}
	data := WeatherData{}

	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return data, err
	}

	defer resp.Body.Close()
	resp_body, _ := io.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	err = json.Unmarshal(resp_body, &data)
	if err != nil {
		fmt.Println("Errored when on unmarshal resp")
		return data, err
	}

	fmt.Println("-----------------------------")
	return data, nil
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]string{"status": "ok"}
	Success(w, http.StatusOK, "Service is healthy", health)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// req := WeatherData{}
	// fmt.Println("-------------------------------")
	// fmt.Println("-------------------------------")
	// json.NewDecoder(r.Body).Decode(&req)

	fmt.Println("-------------------------------")
	city := mux.Vars(r)["city"]

	fmt.Println("city:", city)

	if city == "" {
		Error(w, http.StatusBadRequest, "City is required", nil)

	}

	data, err := query(city)
	if err != nil {
		log.Fatal("[handler] err:", err.Error())
	}
	fmt.Println("-------------------------------")
	Success(w, http.StatusOK, "Request is successful", data)
}
