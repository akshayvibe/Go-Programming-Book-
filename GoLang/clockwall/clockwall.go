package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"net"
	"bufio"
)

type wallClock struct {
	name string
	addr string
}

func main() {
	if len(os.Args)<2 {
		fmt.Fprint(os.Stderr,"usage: clockwall Name=host:port ...")
		os.Exit(1);
	}
	for _,args:=range os.Args[1:]{
		fields:=strings.Split(args, "=")
		if len(fields)!=2{
			log.Fatalf("invalid argument: %s", args)
		}
		go connectToServer(fields[0],fields[1])
	}
	select{}

}
func connectToServer(name,addr string){
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("%s: %v", name, err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Printf("%s: %s\t", name, scanner.Text())
	}
}