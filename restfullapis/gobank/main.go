package main

func main() {
	server := NewAPIServer("localhost:8080")
	server.Run()
}
