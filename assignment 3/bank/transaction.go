package bank

import "fmt"

func Transfer(a *Customer, b *Customer, z int) {
	var s, d int
	if z == 1 {
		fmt.Println("Enter opponent Id:")
		fmt.Scan(&s)
		if b.id == s {
			fmt.Println("Enter amount to transfer")
			fmt.Scan(&d)
			if a.amount > d {
				b.amount = b.amount + d
				a.amount = a.amount - d
				fmt.Println("Transfer Succesfully")
			} else {
				fmt.Println("Insufficient Funds")
			}
		} else {
			fmt.Println("Entered Wrong ID")
		}
	}
	if z == 2 {
		var x, y int
		fmt.Println("Please Enter Your Id")
		fmt.Scan(&x)
		if a.id == x {
			fmt.Println("Enter Money to Withdraw")
			fmt.Scan(&y)
			a.amount = a.amount - y
		}
		if b.id == x {
			fmt.Println("Enter Money to Withdraw")
			fmt.Scan(&y)
			b.amount = b.amount - y
		} else {
			fmt.Println("Entered Wrong Id")
		}

	} else {
		var x, y int
		fmt.Println("Please Enter Your Id")
		fmt.Scan(&x)
		if a.id == x {
			fmt.Println("Enter Money to Deposit")
			fmt.Scan(&y)
			a.amount = a.amount + y
		}
		if b.id == y {
			fmt.Println("Enter Money to Deposit")
			fmt.Scan(&y)
			a.amount = a.amount + y
		} else {
			fmt.Println("Entered Wrong Id")

		}
	}
}
