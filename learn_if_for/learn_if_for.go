package main

import "fmt"

func main() {
	// a := 10
	// if a > 5 {
	// 	fmt.Println("a is bigger than 5")
	// } else {
	// 	fmt.Println("a is less than or equal to 5")
	// }

	// a := 10
	// if a == 5 {
	// 	fmt.Println("a is bigger than 5")
	// } else {
	// 	fmt.Println("a is less than or equal to 5")
	// }

	// a := 10
	// if a > 5 {
	// 	b := a / 2
	// 	fmt.Println("a is bigger than 5", a, b)
	// } else {
	// 	fmt.Println("a is less than or equal to 5")

	// a := 10
	// if b := a / 2; b > 5 {
	// 	fmt.Println("b is bigger than 5:", a, b)
	// } else {
	// 	fmt.Println("b is less than or equal to 5:", a, b)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i > 7 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// a := 3
	// for i := 0; i < 10; i++ {
	// 	if i == a {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Println(i)

	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i > 10 {
	// 		break
	// 	}
	// }
	// fmt.Println(i)

	// s := "Hello, world!"
	// for k, v := range s {
	// 	fmt.Println(k, v, string(v))
	// }

	s := "👋, 🌎"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
