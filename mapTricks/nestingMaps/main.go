package main

func main() {
    m := make(map[string]map[string]int)
    m["fruits"] = make(map[string]int)
    m["fruits"]["apple"] = 1
    m["fruits"]["banana"] = 2
    m["fruits"]["orange"] = 3
    m["vagetables"] = map[string]int{"carrot": 1}
}
