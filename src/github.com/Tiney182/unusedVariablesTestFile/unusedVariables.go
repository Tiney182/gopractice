package main

import "fmt"

func main() {
	var one int // First way of init variable
	_ = one     // _ is used as an ignored variable

	two := 2         // := is  used to initialise in one line
	fmt.Println(two) // system out print

	three := 3  // again one line initialise
	one = three // assigning one to three
}
