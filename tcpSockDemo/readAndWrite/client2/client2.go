package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
    log.Println("begin dial")
    conn, err := net.Dial("tcp", ":8888")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    log.Println("dial finished")

    data := make([]byte, 65536)
    conn.Write([]byte(data))

    time.Sleep(1000 * time.Second)
}
