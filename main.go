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
	
	var x, a, b int 
	a = 2
	b = 4
	x = a^2 + 2*a*b + b^2
	if x == 22 {
		fmt.Println ("true")
		}else{
			fmt.Println ("false")
		}	

	number := 10
	switch number {
	case 1:
		fmt.Println ("number = 10")
	case 10:
		fmt.Println ("number = 10")
	}
		
}
