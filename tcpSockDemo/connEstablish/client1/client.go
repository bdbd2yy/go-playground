package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
    log.Println("begin dial")
    conn, err := net.Dial("tcp", ":8888")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    log.Println("dial ok")
}
