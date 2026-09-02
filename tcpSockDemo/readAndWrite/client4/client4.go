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
    var total int
    for {
        // NOTE: Set a deadline for connection Write
        conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
        n, err := conn.Write(data)
        if err != nil {
            total += n
            log.Printf("write %d bytes, error: %s", n, err)
            break
        }
        total += n
        log.Printf("write %d bytes this time, %d bytes in total", n, total)
    }

    log.Printf("write %d bytes in total", total)
}
