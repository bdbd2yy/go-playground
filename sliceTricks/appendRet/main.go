package main

import (
	"fmt"
	"unsafe"
)

func main() {
    // the address of a never changes. since the slice is a value type
    a := make([]int, 0, 5)
    fmt.Printf("caps: %d, lens: %d\n", cap(a), len(a))
    b := append(a, 1)
    fmt.Printf("&a = %p, data(a) = %p, len(a) = %d, cap(a) = %d\n", &a, unsafe.SliceData(a), len(a), cap(a))
    fmt.Printf("&b = %p, data(b) = %p, len(b) = %d, cap(b) = %d\n", &b, unsafe.SliceData(b), len(b), cap(b))
    fmt.Printf("a: %v, b: %v\n", a, b)
    a = append(a, 2)
    fmt.Printf("a: %v, b: %v\n", a, b)
    fmt.Printf("address of the new a: %p\n", &a)
    c := append(a, 3)
    fmt.Printf("address of the c: %p\n", &c)
    a = c
    fmt.Printf("address of the new a: %p\n", &a)
    fmt.Printf("a: %v, b: %v\n", a, b)
}
