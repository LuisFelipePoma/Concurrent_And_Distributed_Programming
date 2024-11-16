package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

const myIp = "10.142.113.32"

type Info struct {
	Tipo     string
	NodeNum  int
	NodeAddr string
}

type MyInfo struct {
	cont     int
	first    bool
	nextNum  int
	nextAddr string
}

var chMyInfo chan MyInfo
var readyToStart chan bool

var addrs []string
var myNum int

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	myNum = rand.Intn(int(1e6))
	fmt.Println(myNum)
	var n int
	fmt.Print("Ingrese la cantidad de nodos: ")
	fmt.Scanf("%d\n", &n)
	addrs = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Ingrese nodo %d: ", i+1)
		fmt.Scanf("%s\n", &(addrs[i]))
	}
	readyToStart = make(chan bool)
	go func() {
		chMyInfo = make(chan MyInfo)
		chMyInfo <- MyInfo{0, true, int(1e7), ""}
	}()
	go func() {
		gin := bufio.NewReader(os.Stdin)
		fmt.Print("Presione enter para iniciar...")
		gin.ReadString('\n')
		info := Info{"SENDNUM", myNum, myIp}
		for _, addr := range addrs {
			send(addr, info)
		}
	}()

	server()
}
func server() {
	host := fmt.Sprintf("%s:8000", myIp)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	msg, _ := r.ReadString('\n')
	var info Info
	json.Unmarshal([]byte(msg), &info)
	fmt.Println(info)
	switch info.Tipo {
	case "SENDNUM":
		myInfo := <-chMyInfo
		myInfo.cont++
		if info.NodeNum < myNum {
			myInfo.first = false
		} else if info.NodeNum < myInfo.nextNum {
			myInfo.nextNum = info.NodeNum
			myInfo.nextAddr = info.NodeAddr
		}
		go func() {
			chMyInfo <- myInfo
		}()
		if myInfo.cont == len(addrs) {
			if myInfo.first {
				fmt.Println("Soy el primer!! :D")
				criticalSection()
			} else {
				readyToStart <- true
			}
		}
	case "START":
		<-readyToStart
		criticalSection()
	}
}
func criticalSection() {
	fmt.Println("Ha llegado mi turno!! :)")
	myInfo := <-chMyInfo
	if myInfo.nextAddr == "" {
		fmt.Println("I was the last one! :(")
	} else {
		info := Info{Tipo: "START"}
		fmt.Println(myInfo, info)
		send(myInfo.nextAddr, info)
	}
}
func send(remoteAddr string, info Info) {
	remote := fmt.Sprintf("%s:8000", remoteAddr)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	bytesMsg, _ := json.Marshal(info)
	fmt.Fprintln(conn, string(bytesMsg))
}
