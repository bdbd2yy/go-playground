package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
    var wg = sync.WaitGroup{}

    t1 := time.Now()
    // context.WithTimeout(context.Background(), 2*time.Second)
    ctx, _ := context.WithDeadline(context.Background(), t1.Add(8*time.Second))

    wg.Add(1)
    go func() {
        defer wg.Done()
        ip, err := getIp(ctx)
        if err != nil {
            fmt.Println("getIp error:", err)
        }
        fmt.Println("Get ip: ", ip)
    }()

    wg.Wait()
    fmt.Println("Done", time.Since(t1))
}

func getIp(ctx context.Context) (ip string, err error) {
    // listen to the Cancel function
    // simulate the time cost when getting ip
    timer := time.NewTimer(4 * time.Second)
    defer timer.Stop()
    select {
    case <- timer.C:
        return "192.168.200.1", nil
    case <-ctx.Done():
        return "", ctx.Err()
    }
}
