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
var n, min int
var chCont chan int

func main() {
    gin := bufio.NewReader(os.Stdin)
    fmt.Print("Enter port: ")
    port, _ := gin.ReadString('\n')
    port = strings.TrimSpace(port)
    hostname := fmt.Sprintf("10.11.98.229:%s", port)

    fmt.Print("Remote host: ")
    remotehost, _ = gin.ReadString('\n')
    remotehost = strings.TrimSpace(remotehost)

    fmt.Print("N: ")
    port, _ = gin.ReadString('\n')
    port = strings.TrimSpace(port)
    n, _ = strconv.Atoi(strings.TrimSpace(port))
    chCont = make(chan int, 1)
    chCont<- 0

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

    cont := <-chCont
    if cont == 0 {
        min = num;
    } else if num < min {
        send(min)
        min = num
    } else {
        send(num)
    }
    cont++
    if cont == n {
        fmt.Printf("NUMERO FINAL: %d!!\n", min)
        cont = 0
    }
    chCont<- cont
}

func send(num int) {
    conn, _ := net.Dial("tcp", remotehost)
    defer conn.Close()
    fmt.Fprintf(conn, "%d\n", num)
}
