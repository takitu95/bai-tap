package main

import (
	"fmt"
	"strconv"
)

func main() {
	StringResul := PrintTeamNumber(100)
	fmt.Println(StringResul)
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
	return "Team number " + ValueAfterConvert
}
func PrintAddress(Number int, Address string) string {
	ValueAfterConvert1 := strconv.Itoa(Number)
	return "So nha " + ValueAfterConvert1 + ", " + "Dia chi " + Address
}

func PrintDate(Date int) string {
	ValueAfterConvert2 := strconv.Itoa(Date)
	return "Ngay sinh" + ValueAfterConvert2
}
