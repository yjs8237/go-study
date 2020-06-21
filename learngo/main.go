package main

import (
	"fmt"
	"strings"

	"github.com/learngo/banking"
)

type person struct {
	name     string
	age      int
	favorite []string
}

func main() {
	account := banking.MakeAccount("jisang")
	account.ChangeBalance(100)
	fmt.Println(account.CheckBalance())
	err := account.Withdraw(120)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.CheckBalance())
	/*
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

		total := forTest(1, 2, 3, 4, 5)
		fmt.Println(total)

		checkPointer()
		arrayTest()

		food := []string{"kimchi", "ramyun"}
		personJisang := person{name: "jisang", age: 10, favorite: food}
		fmt.Println(personJisang.name)
	*/
}

func arrayTest() {
	arr := []string{"a", "b", "c"}
	arr = append(arr, "d")
	fmt.Println(arr)
}

func checkPointer() {
	a := 2
	b := &a
	*b = 5
	fmt.Println(a, *b)
}

func forTest(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
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
