package main

import "fmt"

func main() {
	var a1, b1 int 
	a1 = 10000
	b1 = 20000
	fmt.Println (a1 + b1)
	
	var (
		name string
		address string
		age int
	)
	name = "Takitu"
	address = "TPHCM"
	age = 27
	fmt.Println (name)
	fmt.Println (address)
	fmt.Println (age)

	var name1, address1, age1 = "Thuan", "Q8", "27"
	fmt.Println (name1)
	fmt.Println (address1)
	fmt.Println (age1)

	var email = "tangthuan1008@gmail.com"
	fmt.Println (email)
	
	var A bool = false
	fmt.Println(A)
}