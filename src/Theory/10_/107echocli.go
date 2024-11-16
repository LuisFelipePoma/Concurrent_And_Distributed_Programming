package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	con, _ := net.Dial("tcp", "10.21.61.155:8000")
	defer con.Close()

	gin := bufio.NewReader(os.Stdin)
	r := bufio.NewReader(con)
	for {
		fmt.Print("Ingrese mensaje: ")
		msg, _ := gin.ReadString('\n')

		fmt.Fprint(con, msg)
		resp, _ := r.ReadString('\n')
		fmt.Printf("Respuesta: %s", resp)
	}
}
