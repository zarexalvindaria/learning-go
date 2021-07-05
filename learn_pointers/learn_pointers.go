package main

import "fmt"

/*
func main() {
	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)
}
*/

/*
// panic caused by nil
func main() {
	var b *int

	fmt.Println(b)
	fmt.Println(*b)
}
*/

/*
// no longer returns a panic
func main() {
	b := new(int)

	fmt.Println(b)
	fmt.Println(*b)
}
*/

/*
// pointer from a function
func setTo10(a *int) {
	*a = 10
}

func main() {
	a := 20
	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)
}
*/

func setTo10Fail(a *int) {
	a = new(int)
	*a = 10
}

func main() {
	a := 20
	fmt.Println(a)
	setTo10Fail(&a)
	fmt.Println(a)
}
