package main

import "fmt"

func main() {
	var s string
	count := 0
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		j := len(s) - 1 - i
		if s[i] != s[j] {
			fmt.Println("Not a Palindrome")
			count = 1
			break
		}
	}
	if count == 0 {
		fmt.Println("Is a Palindrome")
	}

}
