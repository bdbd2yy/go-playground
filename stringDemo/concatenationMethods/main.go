package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
    a := []string{"a", "b", "c"}
    // method 1: +
    ret1 := a[0] + a[1] + a[2]
    // method 2: fmt.Sprintf
    ret2 := fmt.Sprintf("%s%s%s", a[0], a[1], a[2])
    // method 3: strings.Builder
    var sb strings.Builder
    sb.WriteString(a[0])
    sb.WriteString(a[1])
    sb.WriteString(a[2])
    ret3  := sb.String()
    /// method 4: bytes.Buffer
    buf := new(bytes.Buffer)
    buf.WriteString(a[0])
    buf.WriteString(a[1])
    buf.WriteString(a[2])
    ret4 := buf.String()
    // method 5: strings.Join
    ret5 := strings.Join(a, " ")
    fmt.Println(ret1, ret2, ret3, ret4, ret5)
}
