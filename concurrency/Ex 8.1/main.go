package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	"io"
)

func main() {
	port := flag.Int("Port", 8000, "on for the service")
	flag.Parse()
	addr:=fmt.Sprintf("localhost:%d",*port)
	listener,err:=net.Listen("tcp",addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server listening on %s\n", addr)
	for{
		conn,err:=listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConn(conn)
	}
	// select{}
}
func handleConn(c net.Conn){
	defer c.Close()
	for{
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return 
		}
		time.Sleep(1 * time.Second)
	}
}