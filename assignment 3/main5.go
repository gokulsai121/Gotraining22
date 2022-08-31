package main

import (
	"fmt"
	"go-workspace/bank"
)

func main() {
	var z int
	a := bank.RData()
	b := bank.RData1()
	fmt.Println("Press 1 to transfer or Press 2 to Withdraw or Press 3 to Deposit Money")
	fmt.Scan(&z)
	if z == 1 {
		bank.Transfer(a, b, z)
		fmt.Println(a)
		fmt.Println(b)
	}
	if z == 2 {
		bank.Transfer(a, b, z)
		fmt.Println(a)
		fmt.Println(b)
	}
	if z == 3 {
		bank.Transfer(a, b, z)
		fmt.Println(a)
		fmt.Println(b)
	}
}
