// Netcat1 is a read-only TCP client.
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//this will write to the memory 
	go mustCopy(os.Stdout, conn)
	//this will write to the terminal
	mustCopy(conn,os.Stdin)
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
