package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// random number is always one ?
	fmt.Println("I like the number", rand.Intn(10))
	// Println uses standard formatter and as such cannot have format specifiers
	// passed in within the string.
	// %% is a literal %
	fmt.Println("The number will not be placed in the modifier %%g instead placed at the end.", math.Sqrt(7))

	// printf prints via a format specifier.
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
