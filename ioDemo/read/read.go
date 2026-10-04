package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    f, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    buf := make([]byte, 4096)
    for {
        n, err := f.Read(buf)
        // The Go convention is that a reader may report n > 0, err == io.EOF in a single call.
        // So the contract is: process what n gives you first, then look at err.
        fmt.Println(string(buf[:n]))
        if err != nil {
            log.Println(err)
            break
        }
    }

}
