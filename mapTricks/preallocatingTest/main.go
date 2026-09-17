package main

import (
	"fmt"
	"time"
)

func main() {
    m := make(map[int]int)
    start := time.Now()
    cnt := 10000
    for i := range cnt {
        m[i] = i
    }
    fmt.Println(time.Since(start))

    start = time.Now()
    pm := make(map[int]int, cnt)
    for i := range cnt{
        pm[i] = i
    }
    fmt.Println(time.Since(start))
}
