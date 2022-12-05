package main
import("fmt")
func main() {
	println("hello world")
	PrintAPlusB()
}
func PrintAPlusB() {
	var a, b int
	fmt.Print("a=")
	fmt.Scanln(&a)
	fmt.Print("b=")
	fmt.Scanln(&b)
	fmt.Print("a+b=",a+b)
}