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

	enviar(50, ipRemoto)

}
func enviar(num int, ipRemoto string) {

	remotoDir := fmt.Sprintf("%s:%d", ipRemoto, 29002)

	fmt.Println(remotoDir)

	con, _ := net.Dial("tcp", remotoDir)

	defer con.Close()

	fmt.Fprintln(con, string(num))
	fmt.Println(num)
}
