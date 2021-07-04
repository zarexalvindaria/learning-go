package main

import "fmt"

// func main() {
// 	addNumbers(2, 9)
// 	addNumbers(100, -100)
// 	addNumbers(81, 37)
// }

// func addNumbers(a int, b int) {
// 	fmt.Println(a + b)
// }

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleFails(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
	fmt.Println("in doubleFail:", a, arr, s)
}

func main() {
	// a := addNumbers(2, 9)
	// fmt.Println(a)

	// a = addNumbers(4, 10)
	// fmt.Println(a)

	// a = addNumbers(15, -10)
	// fmt.Println(a)

	// div, remainder := divAndRemainder(2, 3)
	// fmt.Println(div, remainder)

	// div, _ = divAndRemainder(10, 4)
	// fmt.Println(div)

	// _, remainder = divAndRemainder(100, -100)
	// fmt.Println(remainder)

	// divAndRemainder(-1, 20)

	a := 1
	arr := [2]int{2, 4}
	s := "hello"
	doubleFails(a, arr, s)
	fmt.Println("in main:", a, arr, s)
}

func addNumbers(a int, b int) int {
	return a + b
}
