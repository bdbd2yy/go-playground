package main

import (
	"fmt"
	"regexp"
)

func main() {
    str := ""

    for {
        fmt.Scanf("%s", &str)
        if str == "q" {
            return
        }
        // determine if the amount returned by the front is correct
        // define the pattern with `` to use escape characters or escape it with a backslash
        pattern := `^[+-]?[0-9]+(\.[0-9]+)?$`
        // there are tow layers of escaping here. we only need to escape the string in go into a valid reexp.
        // pattern := "^[+-]?[0-9]+(\\.[0-9]+)?$" // escape the \

        match, err := regexp.MatchString(pattern, str)
        if err != nil {
            fmt.Printf("err:[%+v]\n", err)
        } else {
            fmt.Printf("match:[%+v]\n", match)
        }
    }

}
