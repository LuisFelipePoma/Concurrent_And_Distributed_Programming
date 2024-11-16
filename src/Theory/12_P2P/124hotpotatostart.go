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
    fmt.Print("Remote host: ")
    remotehost, _ = gin.ReadString('\n')
    remotehost = strings.TrimSpace(remotehost)
    for {
        fmt.Print("Enter number: ")
        str, _ := gin.ReadString('\n')
        num, _ := strconv.Atoi(strings.TrimSpace(str))
        send(num)
    }
}

func send(num int) {
    conn, _ := net.Dial("tcp", remotehost)
    defer conn.Close()
    fmt.Fprintf(conn, "%d\n", num)
}
