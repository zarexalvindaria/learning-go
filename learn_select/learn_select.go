package main

/*
// returns a deadlok
import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int, 1)

	out <- 1

	in <- 2
	fmt.Println("wrote 2 to in")
	v := <-out
	fmt.Println("read", v, "from out") // returns deadlock
}
*/

/*
// using a select statement
import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int, 1)

	out <- 1

	select {
	case in <- 2:
		fmt.Println("wrote 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	}
}
*/

/*
// uses select which randomly selects an input
import "fmt"

func main() {
	in := make(chan int, 1)
	out := make(chan int, 1)

	out <- 1

	select {
	case in <- 2:
		fmt.Println("wrote 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	}
}
*/

/*
// use a default which is always printed
import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	select {
	case in <- 2:
		fmt.Println("wrote 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	default:
		fmt.Println("nothing else works")
	}
}
*/

/*
// use a break statement in a for loop
import "fmt"

func multiples(i int) chan int {
	out := make(chan int)
	curVal := 0
	go func() {
		for {
			out <- curVal * i
			curVal++
		}
	}()
	return out
}

func main() {
	twosCh := multiples(2)
	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}
	// do more stuff
}
*/

/*
// signal to goroutine when it should shutdown -- uses done channel
import (
	"fmt"
	"time"
)

func multiples(i int) (chan int, chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	curVal := 0
	go func() {
		for {
			select {
			case out <- curVal * i:
				curVal++
			case <-done:
				fmt.Println("goroutine shutting down")
				return
			}
		}
	}()
	return out, done
}

func main() {
	twosCh, done := multiples(2)
	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v) // outputs 0 to 20 and "goroutine shutting down"
	}
	close(done)

	//do more stuff
	time.Sleep(1 * time.Second)
}
*/

/*
import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int)
	in2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i := 0; i < 20; i++ {
			in <- i
		}
		close(in)
		wg.Done()
	}()

	go func() {
		for i := 100; i < 110; i++ {
			in2 <- i
		}
		close(in2)
		wg.Done()
	}()

	go func() {
		count := 0
		for count < 2 {
			select {
			case i, ok := <-in:
				if !ok {
					count++
					in = nil
					continue
				}
				fmt.Println("from in, result is", i*2)
			case i, ok := <-in2:
				if !ok {
					count++
					in2 = nil
					continue
				}
				fmt.Println("from in2, result is", i+2)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
*/

/*
// handling every request that comes in -- takes a few seconds to complete in Windows
// but fast in Mac
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	totalStart := time.Now()
	for i := 0; i < 100000; i++ { // initially used 10k then 100k
		start := time.Now()
		wg.Add(1)
		go func(in int) {
			time.Sleep(1 * time.Second)
			out := 2 * in
			fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("total time:", time.Now().Sub(totalStart))
}
*/

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workers := 20000
	pool := make(chan func(int) int, workers)
	for i := 0; i < workers; i++ {
		pool <- func(in int) int {
			time.Sleep(1 * time.Second)
			result := 2 * in
			return result
		}
	}

	var wg sync.WaitGroup
	count := 0
	totalStart := time.Now()
	for i := 0; i < 100000; i++ { // start with 10k then go with 100k
		start := time.Now()
		select {
		case f := <-pool:
			fmt.Println("processing", i)
			count++
			wg.Add(1)
			go func(in int) {
				out := f(in)
				fmt.Println("got", out, "for", in, "after", time.Now().Sub(start))
				pool <- f
				wg.Done()
			}(i)
		default:
			fmt.Println("rejecting request", i, "too busy")
		}
	}
	wg.Wait()
	fmt.Println("total processed:", count)
	fmt.Println("total time:", time.Now().Sub(totalStart))
}
