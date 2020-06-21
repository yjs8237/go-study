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

	// value 를 무시할 수도있다
	length, _ := getLenAndUpper("test")
	fmt.Println(length)

	repeatWord("test", "test1", "test2")
	length, _ = lastProcess("lastFunction")
	fmt.Println(length)
}

func lastProcess(name string) (int, string) {
	defer fmt.Println("finished!! lastProcess")
	return len(name), strings.ToUpper(name)
}

func repeatWord(words ...string) {
	fmt.Println(words)
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
