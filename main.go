package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Hãy nhập số a")
	fmt.Scanln(&a)
	fmt.Println("Hãy nhập số b")
	fmt.Scanln(&b)
	c := a + b
	c1 := a - b
	fmt.Println("a + b =", c)
	fmt.Println("a - b =", c1)
	if a == b {
		fmt.Println("a = b")
	}
	if a > b {
		fmt.Println("a > b")
	}
	if a < b {
		fmt.Println("a < b")
	}
}
