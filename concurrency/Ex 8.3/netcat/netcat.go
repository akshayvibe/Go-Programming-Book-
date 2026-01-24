// Netcat-style TCP client demonstrating synchronization
// using an unbuffered channel (Section 8.4.1).

package main

import (
	"io"
	"log"
	"net"
	"os"
)

/*ANCHOR Exercise 8.3: In netcat3,the interface value conn has the concrete type *net.TCPConn,which
represents a TCP connection. A TCP connection consists of two halves that may be closed
independently using its CloseRead and CloseWrite methods. Modify the main goroutine of
netcat3 to close only the write half of the connection so that the program will continue to
print the final echoes from the reverb1 server even after the standard input has been closed.
 */

func main() {
	// Establish a TCP connection to the server.
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// done is an UNBUFFERED channel used only for synchronization.
	// Its element type is struct{} because no data is needed —
	// the send itself is the event.
	done := make(chan struct{})

	// Start a background goroutine to READ from the network
	// and WRITE to standard output.
	go func() {
		// io.Copy blocks until the connection is closed or an error occurs.
		// This goroutine may still be running after main finishes copying stdin.
		io.Copy(os.Stdout, conn) // errors intentionally ignored

		// This log happens BEFORE the send on done.
		// Due to the happens-before guarantee of unbuffered channels,
		// the main goroutine is guaranteed to see this log before exiting.
		log.Println("done")

		// Send a value on done to signal completion.
		// This send blocks until the main goroutine receives it.
		done <- struct{}{}
	}()

	// Copy from standard input to the network connection.
	// This blocks until stdin is closed (EOF).
	mustCopy(conn, os.Stdin)

	// Close the connection.
	// This causes the server to see EOF and the background goroutine's
	// io.Copy to return.

	
	//NOTE - exercise code
	if tcpcon,ok:=conn.(*net.TCPConn);ok{
		tcpcon.CloseWrite()
	}
	
	// Receive from done.
	// This blocks until the background goroutine signals completion.
	// It establishes a happens-before relationship, guaranteeing
	// that the goroutine has finished before main exits.
	<-done
	conn.Close()	
}

// mustCopy copies data from src to dst and exits the program on error.
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
