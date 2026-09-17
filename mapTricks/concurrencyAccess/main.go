package main

import (
	"fmt"
	"sync"
)

func main() {
    m := make(map[string]int)

    var mu sync.RWMutex

    go func() {
        mu.Lock()
        m["apple"] = 1
        mu.Unlock()
    }()

    go func() {
        mu.RLock()
        // reduce useless operations while holding the lock
        // use a variable to store the value will be better to print the value directly here
        value := m["apple"]
        mu.RUnlock()
        fmt.Println(value)
    }()

    fmt.Scanln()
}
