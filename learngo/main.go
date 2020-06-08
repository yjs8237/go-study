package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "jisang"
	name = "test"
	jisang := "yun"
	jisang = "test"
	fmt.Println(jisang)
	fmt.Println(name)
	fmt.Println(multiply(2, 2))
	len, retName := getLenAndUpper("jisang")
	fmt.Println(len)
	fmt.Println(retName)

	var length int  , var testName string = getLenAndUpper("jisang")
	fmt.Println(length)

}

func getLenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

/*
* 곱셈 해주는 func
 */
func multiply(a int, b int) int {
	return a * b
}

func test() {
	fmt.Printf("Hello test")
}
