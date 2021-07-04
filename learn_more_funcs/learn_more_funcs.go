package main

import "fmt"

// func addOne(a int) int {
// 	return a + 1
// }

// func main() {
// 	myAddOne := addOne
// 	fmt.Println(addOne(1))
// 	fmt.Println(myAddOne(3))
// }

// func main() {
// 	myAddOne := func(a int) int {
// 		return a + 1
// 	}
// 	fmt.Println(myAddOne(3))
// }

// func addOne(a int) int {
// 	return a + 1
// }

// func addTwo(a int) int {
// 	return a + 2
// }

// func printOperation(a int, f func(int) int) {
// 	fmt.Println(f(a))
// }

// func main() {
// 	printOperation(1, addOne)
// 	printOperation(1, addTwo)
// }

/* --
func main() {
	b := 2
	myAddOne := func(a int) int {
		return a + b
	}
	fmt.Println(myAddOne(1))
}
*/

/*
func main() {
	b := 2
	myAddOne := func(a int) {
		b = a + b
	}
	myAddOne(1)
	fmt.Println(b)
}
*/

/*
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}
func main() {
	addOne := makeAdder(1)
	addTwo := makeAdder(2)

	fmt.Println(addOne(1))
	fmt.Println(addTwo(98))
}
*/

func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func main() {
	addOne := makeAdder(1)
	doubleAddOne := makeDoubler(addOne)

	fmt.Println(addOne(1))
	fmt.Println(doubleAddOne(1))
}
