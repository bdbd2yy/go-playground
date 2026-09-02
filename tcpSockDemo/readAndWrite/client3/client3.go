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
        // NOTE: the for loop will continually write data to the client send buffer. After the client send buffer and the server recv buffer are both full, the write operation will be blocked.
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
    time.Sleep(1000 * time.Second)
}
