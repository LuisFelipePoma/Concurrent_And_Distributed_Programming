package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "10.21.61.155:8000")
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handle(con) // podemos atender miles de clienes concurrentemente!
	}
}

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	for {
		msg, _ := r.ReadString('\n')
		fmt.Printf("Recibido: %s", msg)

		fmt.Fprintf(con, msg)
	}
}
