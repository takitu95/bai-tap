package main

import "strconv"

func main(){

 PrintTeamNumber(Number int) string {
	ValueAfterConvert := strconv.Itoa(Number)
	return "Thanh vien nhom" + ValueAfterConvert
}

 PrintAddress(Number int, Address string) string {
	ValueAfterConvert1 := strconv.Itoa(Number)
	return "Dia chi:" + Address + "So nha:" + ValueAfterConvert1
}

 PrintDate(Date int) string {
	ValueAfterConvert2 := strconv.Itoa(Date)
	return "Ngay sinh" + ValueAfterConvert2
}
}