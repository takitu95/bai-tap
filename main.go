package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type AccountRegisterModule struct {
	Username      string
	Password_hash string
	Email         string
}

func main() {
	// Cho A và B tìm tổng và hiệu của A với B, so sánh số A và B
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
	// Kiểm tra đường dẫn của file đang chạy
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	// Kiểm tra số lượng file có trong thư mục
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	// Kiểm tra file có tồn tại trong thư mục hay không
	var filename string
	filename = "main.go"
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Println("File not exist")
	} else {
		fmt.Println("File exist")
	}
	// Nhập username, password, email:
	AccountRegisterObject := AccountRegisterModule{"Adminstrator", "$2y$10$k8P53MG4YkugxUDBzPIc2urVOIMOaS7s3gOjGIoH0JNOUEWwjjka", "Xtown@gmail.com"}
	fmt.Println(AccountRegisterObject)
}
