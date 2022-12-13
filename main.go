package main

import (
	"fmt"
	"strconv"
)

func main() {
	PrintTeamNumber(100)
}

//	func main() {
//		println("hello world")
//		PrintAPlusB()
//	}
func PrintAPlusB() {
	var a, b int
	fmt.Print("a=")
	fmt.Scanln(&a)
	fmt.Print("b=")
	fmt.Scanln(&b)
	fmt.Print("a+b=", a+b)
}

// Define PrintTeamNumber
func PrintTeamNumber(number int) string {
	ValueAfterConvert := strconv.Itoa(number)
	return "Thanh vien nhom" + ValueAfterConvert
}
func PrintAddress(Number int, Address string) string {
	ValueAfterConvert1 := strconv.Itoa(Number)
	return "Dia chi:" + Address + "So nha:" + ValueAfterConvert1
}

func PrintDate(Date int) string {
	ValueAfterConvert2 := strconv.Itoa(Date)
	return "Ngay sinh" + ValueAfterConvert2
}
