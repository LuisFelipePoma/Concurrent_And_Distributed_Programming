// descubrimiento de identificación de cada nodo
// Interfaz de red -> pool de IPs
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Mi IP es ", descubrirIP())
}

func descubrirIP() string {
	var dirIP string = "127.0.0.1"
	interfaces, _ := net.Interfaces()
	for _, valInterface := range interfaces {
		//fmt.Println(valInterface.Name)
		if strings.HasPrefix(valInterface.Name, "Wi-Fi") {
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
