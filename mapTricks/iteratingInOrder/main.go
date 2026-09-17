package main

import (
	"fmt"
	"sort"
)

func main() {
	// You have three common ways to create maps
	// use map literal when you have initail key-value pairs, use make or use make with capacity when you know the approximate size to reduce resizing
	m := map[string]int{"apple": 1, "banana": 2, "orange": 3}
	// the classic way of iterating will not guarantee the order which we already know
	// for key, value := range m {
	//     fmt.Println("key:", key, "value:", value)
	// }
	keys := make([]string, 0, len(m))
	sort.Strings(keys)
	for key := range m {
		keys = append(keys, key)
	}
	for _, key := range keys {
		fmt.Println("key:", key, "value:", m[key])
	}

    // delete and check if the key exist
	delete(m, "orange")

	if value, exist := m["orange"]; exist {
        fmt.Println(value)
	} else {
        fmt.Println("the orange disappeared!")
    }
}
