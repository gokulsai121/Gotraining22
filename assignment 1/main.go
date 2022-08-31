package main

import (
	"fmt"
	"go-workspace/calc"
)

func main() {
	var x int
	var y int
	fmt.Println("Enter Value Of A:")
	fmt.Scan(&x)
	fmt.Println("Enter Value Of B:")
	fmt.Scan(&y)
	fmt.Println("Addition:", x, "+", y, "=", calc.Addition(x, y))
	fmt.Println("Subtraction:", x, "-", y, "=", calc.Subtraction(x, y))
	fmt.Println("Multiplication:", x, "*", y, "=", calc.Multiplication(x, y))
	fmt.Println("Divison:", x, "/", y, "=", calc.Division(x, y))
	fmt.Println("Modulo:", x, "%", y, "=", calc.Modulo(x, y))
}
