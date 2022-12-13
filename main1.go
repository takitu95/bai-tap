package main

import "strconv"

// Define main()
func main() {

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
