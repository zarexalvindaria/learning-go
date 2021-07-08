package main

import (
	"fmt"

	"github.com/zarexalvindaria/language/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen?"))
	// fmt.Println("Hello")
}
