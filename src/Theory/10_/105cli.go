package main

import (
	"fmt"
	"net"
)

// 10.21.61.171
func main() {
	con, _ := net.Dial("tcp", "10.21.61.155:8000")
	defer con.Close()
	fmt.Fprintln(con, "jalados!")
}
