package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

var remotehost string

func main() {
    gin := bufio.NewReader(os.Stdin)
    fmt.Print("Remote port: ")
    port, _ := gin.ReadString('\n')
    port = strings.TrimSpace(port)
    remotehost = fmt.Sprintf("localhost:%s", port)
    send(30)
    send(10)
    send(20)
    send(50)
    send(40)
}

func send(num int) {
    conn, _ := net.Dial("tcp", remotehost)
    defer conn.Close()
    fmt.Fprintf(conn, "%d\n", num)
}

