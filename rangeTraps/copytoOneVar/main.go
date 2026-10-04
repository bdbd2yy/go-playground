package main

import "fmt"

type User struct {
    Name string
}

func main() {
    // Before go1.22, the for range only uses one variable in the whole loop, copy the value into it each time.
    // For example, if you write for _, v := range slice, it equals to
    // for i := 0; i < len(slice); i++ {
    //      v = slice[i]
    // }
    // But after go1.22, every time of for loop will create a new iteration variable, like
    // for i := 0; i < len(slice); i++ {
    //      v := slice[i]
    // }
    nums := []int{1, 2, 3}
    for _, v := range nums {
        v = 100
        fmt.Println(v)
    }
    fmt.Println(nums)
    users := []User{{Name: "bdbd"}, {Name: "ysh"}}
    for _, u := range users {
        u.Name = "ysh"
    }
}
