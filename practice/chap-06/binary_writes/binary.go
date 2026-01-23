package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start_time := time.Now()
	fmt.Println("----------- Download image from web -------------")
	resp, err := http.Get("https://img.freepik.com/free-photo/landscape-morning-fog-mountains-with-hot-air-balloons-sunrise_335224-794.jpg")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create image file
	// image, err := os.Create("bholenath.png")
	image, err := os.Create("baloon.avif")
	if err != nil {
		panic(err)
	}
	defer image.Close()

	// Copy file from internet to file
	n, err := io.Copy(image, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Downloaded %d bytes.", n)
	duration := time.Since(start_time)
	fmt.Println("Total time:", duration)

}
