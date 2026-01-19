package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	port := flag.Int("Port", 8000, "listening on port")
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)
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
		handleConn(conn)

	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)

	for input.Scan() {
		echo(c,input.Text(),1*time.Second)
	}
	c.Close()
}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
