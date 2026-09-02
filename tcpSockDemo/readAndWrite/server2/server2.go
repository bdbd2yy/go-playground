package main

import (
	"fmt"
	"log"
	"net"
	"time"
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
        time.Sleep(10 * time.Second)
        buf := make([]byte, 65536)
        log.Println("start to read from the conn")
        c.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
        n, err := c.Read(buf)
        if err != nil {
            log.Printf("conn read %d, error: %s\n", n, err)
            // err.(net.Error) is a type assertion
            if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
                continue
            }
            return
        }
        log.Printf("read %d bytes, the content is '%s'\n", n, buf[:n])
    }
}
