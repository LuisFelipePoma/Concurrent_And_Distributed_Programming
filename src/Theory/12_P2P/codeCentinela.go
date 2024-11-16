package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Print("Ingrese la IP del nodo remoto: ")
	ipRemoto, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ipRemoto = strings.TrimSpace(ipRemoto)

	enviar(1000, ipRemoto)

}
func enviar(num int, ipRemoto string) {
	remotoDir := fmt.Sprintf("%s:%d", ipRemoto, 9002)
	con, _ := net.Dial("tcp", remotoDir)

	fmt.Fprintln(con, num)
}
