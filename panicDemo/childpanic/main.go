package main

import "fmt"

func child() {
	fmt.Println("child before panic")
	panic("panic")
	fmt.Println("child after panic")
}

func parent() {
	fmt.Println("parent before child")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("child panicked")
		}
	}()
    child()
    fmt.Println("parent after panic")
}

func main() {
    parent()
	fmt.Println("main after parent")
}
