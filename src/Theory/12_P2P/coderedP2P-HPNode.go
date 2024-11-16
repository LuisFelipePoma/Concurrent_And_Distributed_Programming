// armando la red P2P
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

var addrs []string //arreglo: Bitácora de direcciones IP de los miembros de la red
var hostIP string  // Almacenar la IP del nodo

// Servicios
const (
	portRegistrar = 9000 //para registro de un nuevo nodo a la red
	portNotificar = 9001 //para comunicar a todos los nodos de la red que ingresó un nuevo nodo y actualicen su bitácora
	portHP        = 9002
)

func main() {
	//obtener la IP del nodo actual
	hostIP = descubrirIP()
	fmt.Printf("Mi IP es %s\n", hostIP)

	//Rol Servidor, Modo escucha
	//Definir el servicio de Registro
	go servicioRegistrar()
	go servicioHP()

	//Rol Cliente, Modo llamada
	//Definir la solicitud del cliente para unirse a la red
	fmt.Print("Quiero unirme al siguiente nodo con IP: ")
	ipRemoto, _ := bufio.NewReader(os.Stdin).ReadString('\n') //captura la ipremoto desde la consola
	ipRemoto = strings.TrimSpace(ipRemoto)
	//validamos si es el primer nodo de la red
	if ipRemoto != "" {
		//solicitar unirse a la red
		solicitarRegistro(ipRemoto)
	}

	//Rol Servidor, modo escucha
	//Definir el servicio de notificación
	servicioNotificar()
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
func servicioRegistrar() {
	//modo escucha
	//formatear la dirección de escucha
	hostDir := fmt.Sprintf("%s:%d", hostIP, portRegistrar) //192.17.0.2:9000
	ln, _ := net.Listen("tcp", hostDir)
	defer ln.Close()
	for { //modo constante de escucha
		con, _ := ln.Accept()
		//manejar de forma concurrente las conexion
		go handlerRegistro(con)
	}
}
func handlerRegistro(con net.Conn) {
	defer con.Close()
	//1.- capturar lo que devuelve el nodo que acaba de registrarse en la red -> IP del nodo
	ip, _ := bufio.NewReader(con).ReadString('\n')
	ip = strings.TrimSpace(ip)
	//2.- Devolver al nodo nuevo la bitacora de este nodo sin su ip
	//serializacion
	jsonBytes, _ := json.Marshal(addrs)
	fmt.Fprintf(con, "%s\n", string(jsonBytes)) //devuelve la bitácora al nuevo nodo por el mismo puerto y conexion
	//3.- Este nodo comunica al resto de nodos de la red, la IP del nuevo nodo
	notificarTodos(ip)
	//4.- El nodo actual agrega a su bitácora la ip del nuevo nodo
	addrs = append(addrs, ip)
	//log
	fmt.Println(addrs)
}
func notificarTodos(ip string) {
	//recorrer la bitácora y enviar la IP
	for _, valIP := range addrs {
		enviar(valIP, ip)
	}
}
func enviar(valIP, ip string) {
	//modo llamada
	remoteDir := fmt.Sprintf("%s:%d", valIP, portNotificar)
	con, _ := net.Dial("tcp", remoteDir)
	defer con.Close()
	fmt.Fprintln(con, ip) //enviar IP
}
func solicitarRegistro(ipRemoto string) {
	//modo llamada
	remotoDir := fmt.Sprintf("%s:%d", ipRemoto, portRegistrar)
	con, _ := net.Dial("tcp", remotoDir)
	defer con.Close()
	//1.-Nodo envia su IP
	fmt.Fprintln(con, hostIP)
	//2.- Recibir la bitácora que envía el nodo servidor y este nodo debe actualizar su bitácora
	bitacora, _ := bufio.NewReader(con).ReadString('\n')
	bitacora = strings.TrimSpace(bitacora)
	var addrsTemp []string
	json.Unmarshal([]byte(bitacora), &addrsTemp) //deserializar, llevar del formato json a arreglo string
	//3.- actualizar la bitácora
	addrs = append(addrsTemp, ipRemoto) //IPS de la red + Ip nodo servidor al q solicitó unirse a la red
	//log
	fmt.Println(addrs)
}
func servicioNotificar() {
	//modo escucha
	localDir := fmt.Sprintf("%s:%d", hostIP, portNotificar)
	ln, _ := net.Listen("tcp", localDir)
	defer ln.Close()
	for {
		con, _ := ln.Accept()
		go handlerNotificar(con)
	}
}
func handlerNotificar(con net.Conn) {
	defer con.Close()
	//1.-Recibir la IP del nuevo nodo
	ipNewNode, _ := bufio.NewReader(con).ReadString('\n')
	ipNewNode = strings.TrimSpace(ipNewNode)
	//2.- Actualizar la bitácora, agregando esta IP
	addrs = append(addrs, ipNewNode)
	//log
	fmt.Println(addrs)
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

	r := bufio.NewReader(con)
	strNum, _ := r.ReadString('\n')
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
