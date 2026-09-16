package main

import "fmt"

func main() {
    a := make([]int, 5)
    b := make([]int, 5)
    // append a slice
    a = append(a, b...)
    fmt.Println(a)
    b = make([]int, len(a))
    copy(b, a)
    fmt.Println(b)
    // slice copy idiom: if you have to append more elements to b after the copy, maybe these could be more efficient
    // []T(nil) will create a nil slice
    b = append([]int(nil), a...)
    // slice[low:high:max]
    // len = high - low, cap = max - low
    b = append(a[:0:0], a...)
    
    // cut
    // it will overwrite the underlying array
    // O(n)
    a = append(a[:1], a[3:]...)


    // delete
    // eg: delte the element of index 1
    a = []int{1, 2, 3, 4, 5}
    a = append(a[:1], a[2:]...)
    // array: 1, 3, 4, 5, 5 
    fmt.Printf("array: %#v\n", a[:cap(a)])
    fmt.Printf("slice: %#v\n", a)
    // copy returns the number of elements copied
    n := copy(a[1:], a[2:])
    a = a[:1+n]
    // a = a[:i+copy(a[i:], a[i+1:])]
    fmt.Printf("array: %#v\n", a[:cap(a)])
    fmt.Printf("slice: %#v\n", a)

    // delete without preserving the order
    // eg: delete the element of index 2
    a = []int{1, 2, 3, 4, 5}
    a[2] = a[len(a)-1]
    a = a[:len(a)-1]
    fmt.Printf("array: %v\n", a[:cap(a)])
    fmt.Printf("slice: %v\n", a)
    // NOTE: if the type of the element is a pointer or a struct with pointer fileds, attention to the memory leak problem

    // expand
    // eg: insert 3 elements at index 1
    a = append(a[:1], append(make([]int, 3), a[1:]...)...)

    // extend
    // eg: append 10 elements
    a = append(a, make([]int, 10)...)
    fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))

    // extend capacity
    if cap(a) - len(a) < 10 {
        a = append(make([]int, 0, len(a)+10), a...)
    }

    // filter
    n = 0
    for _, x := range a {
        if x > 0 {
            a[n] = x
            n++
        }
    }
    a =  a[:n]

    // insert
    // eg: insert 2 at index 1
    a = []int{1, 3, 4, 5}
    a = append(a[:1], append([]int{2}, a[1:]...)...)
    fmt.Printf("slice: %v\n", a)
    // NOTE: you can avoid creating a new slice(memory grabage) and secong copy using the alternative way
    a = []int{1, 3, 4, 5}
    a = append(a, 0)
    copy(a[2:], a[1:])
    a[1] = 2
    fmt.Printf("slice: %v\n", a)

    // push
    a = append(a, 6)
    a = append([]int{0}, a...)
    fmt.Println(a)
    // pop
    x, a := a[len(a)-1], a[:len(a)-1]
    fmt.Println(x, a)
    x, a = a[0], a[1:]
    fmt.Println(x, a)

}
