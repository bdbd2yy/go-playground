package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
    l, err := net.Listen("tcp", ":8888")
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
        log.Println("accepted a new connection")
        handleConn(c)
    }
}

func handleConn(c net.Conn) {
    defer c.Close()
    for {
        buf := make([]byte, 10)
        // the read funtion will copy the data from the socket receive buffer in kernel to the buffer in the go code
        n, err := c.Read(buf)
        log.Println("start to read from the conn")
        if err != nil {
            fmt.Println(err)
            return
        }
        log.Printf("read %d bytes, the content is '%s'\n", n, buf[:n])
    }
}
