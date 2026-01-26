package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	// // Basic http request not for production
	// resp, err := http.Get("http://localhost:8080/hello")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // resp is not string it is stream so after read we have to close
	// fmt.Println("response body:", resp.Body)

	// io.Copy(os.Stdout, resp.Body)
	// resp.Body.Close()

	// Production need to validate resp status, headers, methods
	url := "http://localhost:8080/hello"

	// Step-1: Create client
	client := &http.Client{Timeout: 10 * time.Second}

	// Step-2: Create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Step-3: Execute request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Reading content length manually
	buf := make([]byte, 1024)
	total := 0
	for {
		n, err := resp.Body.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		total += n
		fmt.Println("data:", string(buf[:n]))
	}

	// --------------- Additional -------------
	fmt.Println("--------------- Additional ------------------")
	fmt.Println("method:", resp.Request.Method)
	fmt.Println("header:", resp.Header)

	contentType := resp.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	if resp.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", resp.ContentLength)
	}

	// // Read data
	// length := 0
	// var buffer [1024]byte
	// r := resp.Body
	// for {
	// 	n, err := r.Read(buffer[0:])
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	length = length + n
	// }
	// fmt.Println("Calculated response data length:", length)
}
