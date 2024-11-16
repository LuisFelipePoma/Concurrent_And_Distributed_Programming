package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
)

const localAddr = "localhost:8003" // su propia IP aquí
const (
	cnum = iota // iota genera valores en secuencia y se reinicia en cada bloque const
	opa
	opb
)

type tmsg struct {
	Code int
	Addr string
	Op   int
}

// Las IP de los demás participantes acá, todos deberían usar el puerto 8000
var addrs = []string{"localhost:8000",
	"localhost:8001",
	"localhost:8002"}

var chInfo chan map[string]int

func main() {
	chInfo = make(chan map[string]int)
	go func() { chInfo <- map[string]int{} }()
	go server()
	time.Sleep(time.Millisecond * 100)
	var op int
	for {
		fmt.Print("Your option: ")
		fmt.Scanf("%d\n", &op)
		msg := tmsg{cnum, localAddr, op}
		for _, addr := range addrs {
			send(addr, msg)
		}
	}
}
func server() {
	if ln, err := net.Listen("tcp", localAddr); err != nil {
		log.Panicln("Can't start listener on", localAddr)
	} else {
		defer ln.Close()
		fmt.Println("Listeing on", localAddr)
		for {
			if conn, err := ln.Accept(); err != nil {
				log.Println("Can't accept", conn.RemoteAddr())
			} else {
				go handle(conn)
			}
		}
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	dec := json.NewDecoder(conn)
	var msg tmsg
	if err := dec.Decode(&msg); err != nil {
		log.Println("Can't decode from", conn.RemoteAddr())
	} else {
		fmt.Println(msg)
		switch msg.Code {
		case cnum:
			concensus(conn, msg)
		}
	}
}
func concensus(conn net.Conn, msg tmsg) {
	info := <-chInfo
	info[msg.Addr] = msg.Op
	if len(info) == len(addrs) {
		ca, cb := 0, 0
		for _, op := range info {
			if op == opa {
				ca++
			} else {
				cb++
			}
		}
		if ca > cb {
			fmt.Println("GO A!")
		} else {
			fmt.Println("GO B!")
		}
		info = map[string]int{}
	}
	go func() { chInfo <- info }()
}
func send(remoteAddr string, msg tmsg) {
	if conn, err := net.Dial("tcp", remoteAddr); err != nil {
		log.Println("Can't dial", remoteAddr)
	} else {
		defer conn.Close()
		fmt.Println("Sending to", remoteAddr)
		enc := json.NewEncoder(conn)
		enc.Encode(msg)
	}
}
