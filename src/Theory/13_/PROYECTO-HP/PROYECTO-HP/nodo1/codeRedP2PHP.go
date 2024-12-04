// armando la red P2P
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

var addrs []string //arreglo: Bitácora de direcciones IP de los miembros de la red
var hostIP string  // Almacenar la IP del nodo

// Servicios
const (
	portHP = 9002 //Servicio HP
)

func main() {
	//obtener la IP del nodo actual
	hostIP = descubrirIP()
	fmt.Printf("Mi IP es %s\n", hostIP)

	addrs = []string{
		"172.20.0.3", "172.20.0.4"}

	//Rol Servidor, Modo escucha
	servicioHP()
}

func descubrirIP() string {
	var dirIP string = "127.0.0.1"
	interfaces, _ := net.Interfaces()
	for _, valInterface := range interfaces {
		//fmt.Println(valInterface.Name)
		if strings.HasPrefix(valInterface.Name, "eth0") {
			direcciones, _ := valInterface.Addrs()
			for _, valDireccion := range direcciones {
				switch d := valDireccion.(type) {
				case *net.IPNet:
					if d.IP.To4() != nil {
						//fmt.Println(d.IP.To4())
						dirIP = d.IP.String()
					}
				}

				//fmt.Println(valDireccion.String())
			}
		}
	}
	return dirIP
}
func servicioHP() {
	//modo escucha
	localDir := fmt.Sprintf("%s:%d", hostIP, portHP)

	ln, _ := net.Listen("tcp", localDir)
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handlerHP(con)
	}
}
func handlerHP(con net.Conn) {
	defer con.Close()

	strNum, _ := bufio.NewReader(con).ReadString('\n')
	fmt.Println(strNum)
	num, _ := strconv.Atoi(strings.TrimSpace(strNum))
	fmt.Printf("Recibimos el %d\n", num)
	if num == 0 {
		fmt.Println("Perdimos!!!!!")
	} else {
		enviarHP(num - 1)
	}
}
func enviarHP(num int) {
	idx := rand.Intn(len(addrs)) //obtener el indice del próximo IP a enviar el número
	fmt.Println(idx)
	fmt.Printf("Enviando %d a %s\n", num, addrs[idx])
	remoteDir := fmt.Sprintf("%s:%d", addrs[idx], portHP)
	conn, _ := net.Dial("tcp", remoteDir)
	defer conn.Close()
	fmt.Fprintln(conn, num)
}
