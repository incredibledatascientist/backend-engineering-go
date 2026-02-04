package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

// go get golang.org/x/net/websocket
// go get github.com/gorilla/websocket@latest
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func healthHander(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)

	w.Write([]byte(`{"status":"ok"}`))
}

func rootHander(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	fmt.Fprintf(w, "Welcome!\n")
	fmt.Fprintf(w, "Please use /ws for WebSocket!")
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf()
	// You can't use these http methods.

	log.Println("Serving:", r.URL.Path, "from", r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	for {
		// reading from the websocket.
		msgType, msg, err := ws.ReadMessage()

		if err != nil {
			log.Println("From", r.Host, "read", err)
			break
		}
		message := string(msg)
		fmt.Println("<--", message, " & type:", msgType)

		// writing to websocket.
		if strings.ToLower(message) == "ping" {
			err = ws.WriteMessage(websocket.TextMessage, []byte("pong"))
		} else {
			err = ws.WriteMessage(msgType, msg)
		}

		// ------- Message Type ---------
		// websocket.TextMessage    // (Text / JSON)
		// websocket.BinaryMessage // (FILES & RAW DATA)
		// websocket.CloseMessage // (GOODBYE PROPERLY)
		// websocket.PingMessage //
		// websocket.PongMessage
		// err = ws.WriteMessage(websocket.TextMessage, resp)
		// err = ws.WriteMessage(websocket.PongMessage, resp)
		if err != nil {
			fmt.Print("--> Error:", err)
			break
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(rootHander))
	mux.Handle("/ws", http.HandlerFunc(websocketHandler))
	mux.Handle("/health", http.HandlerFunc(healthHander))

	server := &http.Server{
		Addr:         "localhost:80",
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("Server is listening on addr:", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
