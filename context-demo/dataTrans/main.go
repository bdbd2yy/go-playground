package main

import (
	"context"
	"fmt"
)

func main() {
    c := context.Background()
    c = context.WithValue(c, "name", "bdbd")
    getUser(c)
}

func getUser(ctx context.Context) {
    fmt.Println(ctx.Value("name"))
}
