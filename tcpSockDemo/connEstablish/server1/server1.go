package main

import (
	"fmt"
	"net"
)

func main() {
    l, err := net.Listen("tcp", ":8888")
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        con, err := l.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
        go handlecon(con)
    }
}

func handlecon(c net.Conn) {
    defer c.Close()
    for {
        // read from the connection
        // write to the connection
    }
}
