package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func 来自打工人的问候() {
	问候语 := "早安，打工人　☺"
	fmt.Println(问候语)
	bytes := []byte(问候语)
	fmt.Println(hex.EncodeToString(bytes))
}

func TestChar() {
	fmt.Println(len("早"))
	fmt.Println(len([]byte("早")))
	fmt.Println(len([]rune("早")))
}

func main() {
	// base64 to string
	base64Bytes := []byte("hello golang base64 编码快乐")
	base64String := base64.StdEncoding.EncodeToString(base64Bytes)
	fmt.Println(base64String)
	// string to base64
	base64By, _ := base64.StdEncoding.DecodeString(base64String)
	fmt.Println(base64By)

	来自打工人的问候()
	TestChar()

}
