package main

import (
	"fmt"
)
func main() {
	var a = []int{91, 42, 23, 14, 15, 76, 87, 28, 19, 95}
	o := make(chan int)
	e := make(chan int)

	go odd(o)
	go even(e)

	for i := 0; i < len(a); i++ {
		if a[i]%2 != 0 {
			o <- a[i]
		} else {
			e <- a[i]
		}
	}

}

func odd(ch <-chan int) {
	for v := range ch {
		fmt.Println("odd:", v)
	}
}
func even(ch <-chan int) {
	for v := range ch {
		fmt.Println("even:", v)
	}
}
