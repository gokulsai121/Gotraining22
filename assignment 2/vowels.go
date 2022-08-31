package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		v := string(s[i])
		if v == "a" || v == "e" || v == "i" || v == "o" || v == "u" {
			s = strings.ReplaceAll(s, v, "")
			i--
		}
	}
	fmt.Println(s)
}
