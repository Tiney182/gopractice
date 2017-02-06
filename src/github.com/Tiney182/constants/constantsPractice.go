package main

import "fmt"

// constants are defined with the const keyword
const Pi = 3.14

// constants cannot be insansiated via :=
// for example Pi := 3.14
// the general naming convention for constants is camelcasing
// this is upper or lower depending on if you want the to export the constant
// Constants can be character, string, boolean, or numeric values.

func main() {
	fmt.Println(Pi)
}
