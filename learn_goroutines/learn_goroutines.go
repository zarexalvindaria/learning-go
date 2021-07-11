package main

/*
import "fmt"

func runMe() {
	fmt.Println("Hello from a goroutine")
}

func main() {
	go runMe() // does not return anything
}
*/

/*
import (
	"fmt"
	"time"
)

func runMe() {
	fmt.Println("Hello from a goroutine")
}

func main() {
	go runMe()
	time.Sleep(1 * time.Second) // workaround but not a good practice
}
*/

/*
import (
	"fmt"
	"sync"
)

func runMe() {
	fmt.Println("Hello from a goroutine")
}

func main() {
	var wg sync.WaitGroup // best practice when using a goroutine
	wg.Add(1)	// is to use a waitgroup as it starts and ends immediately
	go func() {
		runMe()
		wg.Done()
	}()

	wg.Wait()
}
*/

/*
// pass data to a goroutine
import (
	"fmt"
	"sync"
)

func runMe(name string) {
	fmt.Println("Hello to", name, "from a goroutine")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(name string) {
		runMe(name)
		wg.Done()
	}("Bob")

	wg.Wait()
}
*/

/*
// incorrect way to pass param to goroutine
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i) // returns several 10s and a few different numbers
			// which is not expected
			wg.Done()
		}()
	}
	wg.Wait()
}
*/

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(localI int) {
			fmt.Println(localI)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
