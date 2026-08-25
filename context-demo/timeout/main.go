package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
    var wg sync.WaitGroup
    ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)
    wg.Add(1)
    t1 := time.Now()
    go func() {
        defer wg.Done()
        ip, err := getIp(ctx)
        if err != nil {
            fmt.Println("getIp error", err)
        }
        fmt.Println(ip)
    }()
    wg.Wait()
    fmt.Println("Done", time.Since(t1))
}

func getIp(c context.Context) (string, error) {
    timer := time.NewTimer(2*time.Second)
    select {
    case <- timer.C:
        return "192.168.200.1", nil
    case <- c.Done():
        return "", c.Err()
    }
}
