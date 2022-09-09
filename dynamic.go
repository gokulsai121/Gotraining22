package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	fmt.Println("Enter week Number from 1 - 7")
	fmt.Scan(&a)
	switch a {
	case 1:
		z := "Monday"
		fmt.Println("Type:", reflect.TypeOf(z), "\n", "Value:", reflect.ValueOf(z))
	case 2:
		z := "Tuesday"
		fmt.Println("Type:", reflect.TypeOf(z), "\n", "Value:", reflect.ValueOf(z))
	case 3:
		z := "Wednesday"
		fmt.Println("Type:", reflect.TypeOf(z), "\n", "Value:", reflect.ValueOf(z))
	case 4:
		z := "Thursday"
		fmt.Println(z)
		fmt.Println("Type:", reflect.TypeOf(z), "\n", "Value:", reflect.ValueOf(z))
	case 5:
		z := "Friday"
		fmt.Println("Type:", reflect.TypeOf(z), "\n", "Value:", reflect.ValueOf(z))
	case 6:
		z := "Saturday"
		fmt.Println("Type:", reflect.TypeOf(z), "\n", "Value:", reflect.ValueOf(z))
	case 7:
		z := "Sunday"
		fmt.Println("Type:", reflect.TypeOf(z), "\n", "Value:", reflect.ValueOf(z))
	}

}
