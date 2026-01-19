package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	// "log/slog"
	"net"
	"strings"
	"time"
)

func main() {
	port := flag.Int("Port", 8000, "listening on port")
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)
	log.Println("Server is starting...")
	listner, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Println("Client connected to",conn.RemoteAddr())
		go handleConn(conn)

		
	}

}
func handleConn(c net.Conn) {
	// defer 
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		 go echo(c,input.Text(),1*time.Second)
	}
	log.Println("Client disconnected to",c.RemoteAddr())
}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
