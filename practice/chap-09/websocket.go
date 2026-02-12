package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

// go get golang.org/x/net/websocket
// go get github.com/gorilla/websocket@latest
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		if origin == "" {
			return true
		}
		return strings.Contains(origin, r.Host)
	},
}

func healthHander(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
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
		log.Println("upgrade error:", err)
		return
	}
	defer ws.Close()

	const (
		maxMessageSize = 512
		pongWait       = 60 * time.Second
		pingPeriod     = (pongWait * 9) / 10
		writeWait      = 10 * time.Second
	)

	ws.SetReadLimit(maxMessageSize)
	_ = ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error {
		return ws.SetReadDeadline(time.Now().Add(pongWait))
	})

	pingTicker := time.NewTicker(pingPeriod)
	defer pingTicker.Stop()

	go func() {
		for range pingTicker.C {
			_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}()

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
		_ = ws.SetWriteDeadline(time.Now().Add(writeWait))
		if strings.EqualFold(message, "ping") {
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

	serverErr := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-stop:
		log.Println("received signal:", sig)
	case err := <-serverErr:
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("graceful shutdown failed:", err)
	}
}
