package main

import (
	"fmt"
)

func IsPrime(n int32) bool {
	fmt.Scanln(&n)
	if n <= 1 {
		return false
	}
	for i := int32(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	var n int32
	fmt.Println(IsPrime(n))
}
