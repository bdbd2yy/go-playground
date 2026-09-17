package main

import "fmt"

type User struct {
    age int
    cute bool
}

func main() {
    // When you use struct as a map's key, the fields in the struct must be comparable. or it will cause a compile-time error
    // you may be confused: why can't i add the value to the key's struct?
    // because you are holding the key's struct to query the name of of the User, if you know the name, why do you come here use the struct to query?
    // this is similar to querying a database using a composite primary key
    m := make(map[User]string)
    m[User{3, true}] = "ysh"
    m[User{20, false}] = "ysh" 

    fmt.Println("Who is the cute baby? ", m[User{3, true}])
}
