// Clock1 is a TCP server that periodically writes the time.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
    // 1. Define the flag
    port := flag.Int("port", 8000, "port to listen on")
    
    // 2. Parse the flags (IMPORTANT!)
    flag.Parse()
    
    // 3. Use a colon (:) between address and port
    address := fmt.Sprintf("localhost:%d", *port)
    
    listener, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Server listening on %s\n", address)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go handleConn(conn)
    }
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
