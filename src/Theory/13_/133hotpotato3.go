package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

var addrs []string

func main() {
	myip := "10.142.113.32"
	fmt.Printf("Soy %s\n", myip)
	go registerServer(myip)
	go hotServer(myip)

	gin := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese direccion remota: ")
	remoteIp, _ := gin.ReadString('\n')
	remoteIp = strings.TrimSpace(remoteIp)

	if remoteIp != "" {
		registerSend(remoteIp, myip)
	}

	go func() {
		fmt.Print("Ingrese num: ")
		strNum, _ := gin.ReadString('\n')
		if strNum != "" {
			num, _ := strconv.Atoi(strings.TrimSpace(strNum))
			hotSend(num)
		}
	}()

	notifyServer(myip)
}
func hotServer(hostAddr string) {
	host := fmt.Sprintf("%s:8002", hostAddr)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleHot(conn)
	}
}
func handleHot(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	strNum, _ := r.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(strNum))
	fmt.Printf("Recibimos el %d\n", num)
	if num == 0 {
		fmt.Println("Perdimos")
	} else {
		hotSend(num - 1)
	}
}
func hotSend(num int) {
	idx := rand.Intn(len(addrs))
	fmt.Println(idx)
	fmt.Printf("Enviando %d a %s\n", num, addrs[idx])
	remote := fmt.Sprintf("%s:8002", addrs[idx])
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintln(conn, num)
}
func registerSend(remoteAddr, hostAddr string) {
	remote := fmt.Sprintf("%s:8000", remoteAddr)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()

	// Enviar direccion
	fmt.Fprintln(conn, hostAddr)

	// Recibir lista de direcciones
	r := bufio.NewReader(conn)
	strAddrs, _ := r.ReadString('\n')
	var respAddrs []string
	json.Unmarshal([]byte(strAddrs), &respAddrs)

	// agregamos direcciones de nodos a propia libreta
	for _, addr := range respAddrs {
		if addr == remoteAddr {
			return
		}
	}
	addrs = append(respAddrs, remoteAddr)
	fmt.Println(addrs)
}
func registerServer(hostAddr string) {
	host := fmt.Sprintf("%s:8000", hostAddr)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleRegister(conn)
	}
}
func handleRegister(conn net.Conn) {
	defer conn.Close()

	// Recibimos addr del nuevo nodo
	r := bufio.NewReader(conn)
	remoteIp, _ := r.ReadString('\n')
	remoteIp = strings.TrimSpace(remoteIp)

	// respondemos enviando lista de direcciones de nodos actuales
	byteAddrs, _ := json.Marshal(addrs)
	fmt.Fprintf(conn, "%s\n", string(byteAddrs))

	// notificar a nodos actuales de llegada de nuevo nodo
	for _, addr := range addrs {
		notifySend(addr, remoteIp)
	}

	// Agregamos nuevo nodo a la lista de direcciones
	for _, addr := range addrs {
		if addr == remoteIp {
			return
		}
	}
	addrs = append(addrs, remoteIp)
	fmt.Println(addrs)
}
func notifySend(addr, remoteIp string) {
	remote := fmt.Sprintf("%s:8001", addr)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintln(conn, remoteIp)
}
func notifyServer(hostAddr string) {
	host := fmt.Sprintf("%s:8001", hostAddr)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleNotify(conn)
	}
}
func handleNotify(conn net.Conn) {
	defer conn.Close()

	// Recibimos addr del nuevo nodo
	r := bufio.NewReader(conn)
	remoteIp, _ := r.ReadString('\n')
	remoteIp = strings.TrimSpace(remoteIp)

	// Agregamos nuevo nodo a la lista de direcciones
	for _, addr := range addrs {
		if addr == remoteIp {
			return
		}
	}
	addrs = append(addrs, remoteIp)
	fmt.Println(addrs)
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
