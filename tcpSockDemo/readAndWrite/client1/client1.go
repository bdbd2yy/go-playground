package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("lack of the second arg")
        return
    }
    log.Println("begin dial")
    conn, err := net.Dial("tcp", ":8888")
    if err != nil {
        fmt.Println(err)
        return
    }
    log.Println("dial finished")
    defer conn.Close()
    time.Sleep(2 * time.Second)

    data := os.Args[1]
    conn.Write([]byte(data))

    time.Sleep(1000 * time.Second)
}
