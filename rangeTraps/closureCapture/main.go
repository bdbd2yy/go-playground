// you can find the tutorial here: https://go.dev/doc/faq#closures_and_goroutines
// This issue has been fixed in go1.22 and later versions
package main

import "fmt"

func main() {
    done := make(chan struct{})

    values := []string{"a", "b", "c"}

    // it may print: c c c
    // Because the closure in go actually capture the variable itself instead of the copy of it
    // Here the closure captures the same variable v
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- struct{}{}
        }()
    }
    // So the fix to it is also simple. Let the closure use the copy
    for _, v := range values {
        // Directly copy the value in the for loop using variable shadowing
        v := v
        go func() {
            fmt.Println(v)
            done <- struct{}{}
        }()
    }
    // Even though the issue has been fixed, but passing the copy into the closure is still a good coding style
    for _, v := range values {
        // This will pass the copy of the current v into the closure
        go func(v string) {
            fmt.Println(v)
            done <- struct{}{}
        }(v)
    }
    for _ = range values {
        <- done
    }
}
