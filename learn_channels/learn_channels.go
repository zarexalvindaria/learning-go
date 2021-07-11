package main

/*
import "fmt"

func main() {
	in := make(chan string)
	out := make(chan string)
	go func() {
		name := <-in
		out <- fmt.Sprintf("Hello," + name)
	}()
	in <- "Bob"
	close(in)
	message := <-out
	fmt.Println(message)
}
*/

/*
// using a buffer channel
import "fmt"

func main() {
	out := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(localI int) {
			out <- localI * 2
		}(i)
	}
	var result []int
	for i := 0; i < 10; i++ {
		val := <-out
		result = append(result, val)
	}
	fmt.Println(result)
}
*/

/*
import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	in <- 1
	in <- 2
	o1 := <-out
	o2 := <-out
	fmt.Println(o1, o2) // deadlock! all goroutines are asleep
}
*/

/*
// returns correctly
import "fmt"

func main() {
	in := make(chan int, 2) // change from previous program
	out := make(chan int, 2) // change from previous program

	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	in <- 1
	in <- 2
	o1 := <-out
	o2 := <-out
	fmt.Println(o1, o2) // returns 2, 4
}
*/

/*
// does not need a goroutine
import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for {
			i := <-in
			out <- i * 2
		}
	}()

	in <- 1
	o1 := <-out
	in <- 2
	o2 := <-out
	fmt.Println(o1, o2) // returns 2, 4
}
*/

import "fmt"

func main() {
	in := make(chan int, 10)
	out := make(chan int)

	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)

	go func() {
		for {
			i, ok := <-in
			if !ok {
				close(out)
				break
			}
			out <- i * 2
		}
	}()
	for v := range out {
		fmt.Println(v)
	}
}
