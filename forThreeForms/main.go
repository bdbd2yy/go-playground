package main

import (
	"fmt"
)

func main() {
    // while form for loop
    loveme := true
    for !loveme {
        fmt.Println("if you don't love me any more, love loop ends")
    }
    // for range loop
    pair := map[string]string{"ysh": "unknown"}
    for boyfriend, grilfriend := range pair {
        fmt.Printf("%s and %s are a pair", boyfriend, grilfriend)
    }
    for v := range 10 {
        fmt.Printf("This is the %d time i miss you", v)
    }
    for range 10 {
        fmt.Println("This is the i don't know how many times i miss you")
    }
    // traditional for loop
    for i := 0; i < 10; i++ {
        fmt.Println("This is the classic for loop")
    }
    for {
        fmt.Println("This is a infinite loop")  
    }
}
