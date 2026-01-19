package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"io"
	"os"
	"strings"
	"bufio"
	"path/filepath"
)

func main() {
	port := flag.Int("Port",8021,"is listenning")
	flag.Parse()
	addrr:=fmt.Sprintf("localhost:%d",*port)
	listener,err:=net.Listen("tcp",addrr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("FTP Server started on localhost:8021")
	for{
		conn,err:=listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleFTPConn(conn) // Handle each client concurrently
	}
}
func handleFTPConn(c net.Conn) {
	defer c.Close()
	// Initial working directory is where the server started
	cwd, _ := os.Getwd()
	fmt.Fprintf(c, "Connected to GoFTP. Working Dir: %s\n", cwd)

	input := bufio.NewScanner(c)
	for input.Scan() {
		args := strings.Fields(input.Text())
		if len(args) == 0 {
			continue
		}

		command := args[0]
		switch command {
		case "ls":
			listDir(c, cwd)
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(c, "Usage: cd <dir>")
				continue
			}
			newDir := filepath.Join(cwd, args[1])
			if info, err := os.Stat(newDir); err == nil && info.IsDir() {
				cwd = newDir
				fmt.Fprintf(c, "OK: CWD is now %s\n", cwd)
			} else {
				fmt.Fprintf(c, "Error: %v\n", err)
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintln(c, "Usage: get <filename>")
				continue
			}
			sendFile(c, filepath.Join(cwd, args[1]))
		case "close":
			fmt.Fprintln(c, "Goodbye!")
			return
		default:
			fmt.Fprintf(c, "Unknown command: %s\n", command)
		}
		fmt.Fprint(c, "> ") // Command prompt for the client
	}
}

func listDir(w io.Writer, dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}
	for _, file := range files {
		fmt.Fprintf(w, "%s\t", file.Name())
	}
	fmt.Fprintln(w)
}

func sendFile(w io.Writer, path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(w, "Error: %v\n", err)
		return
	}
	defer file.Close()
	io.Copy(w, file)
	fmt.Fprintln(w, "\n--- End of File ---")
}