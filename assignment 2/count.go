package main

import "fmt"


func main() {
	alphabets := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0,
		"g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0,
		"n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0,
		"u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0}
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		count := 0
		v := string(s[i])
		for j := 0; j < len(s); j++ {
			h := string(s[j])
			if v == h {
				count += 1
			}
		}
		alphabets[v] = count
	}
	fmt.Println(alphabets)
}
