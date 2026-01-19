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
// 1. READS from the network (conn) and WRITES to your screen (os.Stdout)
go mustCopy(os.Stdout, conn) 

// 2. READS from your keyboard (os.Stdin) and WRITES to the network (conn)
mustCopy(conn, os.Stdin)	
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
