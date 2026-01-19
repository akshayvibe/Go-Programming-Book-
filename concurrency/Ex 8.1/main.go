package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
	"io"
)

/*ANCHOR - Exercis e 8.1: Modify clock2 to accepta port number, and write a program, clockwall, that
ac ts as a client of several clo ck ser vers at once, reading the times fro m each one and displaying
the results in a table, akin to the wall of clocks seen in some business offices. If you have
access to geographic ally distr ibute d computers, run instances remotely ; other wise run local
inst ances on dif ferent ports with fake time zones.
$ TZ=US/Eastern ./clock2 -port 8010 &
$ TZ=Asia/Tokyo ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
$ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
*/

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