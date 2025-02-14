package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		conn, _ := net.Dial("tcp", "localhost:8080")
		fmt.Fprintf(conn, text+"\n")
		time.Sleep(2 * time.Millisecond)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf(message)
		conn.Close()
	}
}
