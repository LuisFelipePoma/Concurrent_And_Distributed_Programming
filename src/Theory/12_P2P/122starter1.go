package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "net"
    "os"
    "strings"
    "time"
)

var remotehost string

func main() {
    gin := bufio.NewReader(os.Stdin)
    fmt.Print("Remote host: ")
    remotehost, _ = gin.ReadString('\n')
    remotehost = strings.TrimSpace(remotehost)
    rand.Seed(time.Now().UTC().UnixNano())
    for i := 0; i < 19; i++ {
        send(rand.Intn(900) + 100)
    }
}

func send(num int) {
    conn, _ := net.Dial("tcp", remotehost)
    defer conn.Close()
    fmt.Fprintf(conn, "%d\n", num)
}

