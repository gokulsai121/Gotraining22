package main

import "fmt"

func main() {
	arr := [3]int{1, 1, 2}
	arr1 := [len(arr)]int{}
	i := 1
	arr1[0] = arr[0]
	for j := 0; j < len(arr); j++ { //0,1,2
		if arr[j] != arr[i] {
			arr1[i] = arr[j]
			i++
		}

	}
	fmt.Println(arr1)
}
