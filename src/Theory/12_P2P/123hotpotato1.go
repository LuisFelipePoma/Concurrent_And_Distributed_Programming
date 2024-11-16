package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
    "strconv"
)

var remotehost string

func main() {
    gin := bufio.NewReader(os.Stdin)
    fmt.Print("Enter port: ")
    port, _ := gin.ReadString('\n')
    port = strings.TrimSpace(port)
    hostname := fmt.Sprintf("10.11.98.229:%s", port)

    fmt.Print("Remote host: ")
    remotehost, _ = gin.ReadString('\n')
    remotehost = strings.TrimSpace(remotehost)

    // Listener!
    ln, _ := net.Listen("tcp", hostname)
    defer ln.Close()
    for {
        conn, _ := ln.Accept()
        go handle(conn)
    }
}

func handle(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    str, _ := r.ReadString('\n')
    num, _ := strconv.Atoi(strings.TrimSpace(str))
    fmt.Printf("Nos ha llegado el %d\n", num)
    if num == 0 {
        fmt.Println("Perdimos! :(")
    } else {
        send(num - 1)
    }
}

func send(num int) {
    conn, _ := net.Dial("tcp", remotehost)
    defer conn.Close()
    fmt.Fprintf(conn, "%d\n", num)
}
