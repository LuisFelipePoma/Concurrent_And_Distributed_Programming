package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "10.21.61.155:8000")
	defer ln.Close()
	con, _ := ln.Accept()
	defer con.Close()
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	fmt.Print(msg)
}
