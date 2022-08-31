package bank

type Customer struct {
	name   string
	email  string
	phone  int
	id     int
	accno  int
	amount int
}

func RData() *Customer {
	a := &Customer{
		"abc",
		"abc@e.com",
		9823456456,
		123,
		123456,
		122,
	}
	return a
}
func RData1() *Customer {
	b := &Customer{
		"xyz",
		"xyz@e.com",
		9823456456,
		124,
		11200,
		125,
	}
	return b
}
