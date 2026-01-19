package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 1. Check if the user provided an address (e.g., localhost:8021)
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run netcat.go host:port")
	}
	address := os.Args[1]

	// 2. Connect to the server
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	
	// 3. Close the connection when the program exits
	defer conn.Close()

	// 4. Create a channel to wait for the server to finish
	done := make(chan struct{})

	// 5. Background Goroutine: Read from server and print to terminal
	go func() {
		io.Copy(os.Stdout, conn) // This blocks until the server stops sending
		log.Println("Server disconnected")
		done <- struct{}{} // Tell the main thread we are finished
	}()

	// 6. Main Goroutine: Read from terminal (keyboard) and send to server
	// This allows you to type commands like "ls", "cd", or "get"
	mustCopy(conn, os.Stdin)

	// Wait for the background reader to finish before closing
	<-done 
}

// mustCopy is a helper function commonly used in GOPL
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}