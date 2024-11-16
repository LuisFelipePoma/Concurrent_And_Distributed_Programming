package main

import (
	"bufio"
	"encoding/json"
    "fmt"
    "net"
    "os"
    "strings"
)

var addrs []string

var hostaddr string

const (
	registerport = 8000
	notifyport = 8001
)

func handleNotify(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	ip, _ := r.ReadString('\n')
	ip = strings.TrimSpace(ip)
	addrs = append(addrs, ip)
}
func notifyServer() {
	hostname := fmt.Sprintf("%s:%d", hostaddr, notifyport)
	ln, _ := net.Listen("tcp", hostname)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleNotify(conn)
	}
}
func notify(addr, ip string) {
	remote := fmt.Sprintf("%s:%d", addr, notifyport)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", ip)
}
func tellEverybody(ip string) {
	for _, addr := range addrs {
		notify(addr, ip)
	}
}
func handleRegister(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	ip, _ := r.ReadString('\n')
	ip = strings.TrimSpace(ip)
	bytes, _ := json.Marshal(addrs)
	fmt.Fprintf(conn, "%s\n", string(bytes))
	tellEverybody(ip)
	addrs = append(addrs, ip)
}
func registerServer() {
	hostname := fmt.Sprintf("%s:%d", hostaddr, registerport)
	ln, _ := net.Listen("tcp", hostname)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleRegister(conn)
	}
}
func registerClient(remoteAddr string) {
	remote := fmt.Sprintf("%s:%d", remoteAddr, registerport)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", hostaddr)
	r := bufio.NewReader(conn)
	msg, _ := r.ReadString('\n')
	var respAddrs []string
	json.Unmarshal([]byte(msg), &respAddrs)
	addrs = append(respAddrs, remoteAddr)
	fmt.Println(addrs)
}

func main() {
	hostaddr = myIp()
	go registerServer()
	gin := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese ip remota: ")
	remoteAddr, _ := gin.ReadString('\n')
	remoteAddr = strings.TrimSpace(remoteAddr)
	if remoteAddr != "" {
		registerClient(remoteAddr)
	}
	notifyServer()
}
func myIp() string { // mandrakeando ando
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "Local") {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					return v.IP.String()
				case *net.IPAddr:
					return v.IP.String()
				}
			}
		}
	}
	return ""
}