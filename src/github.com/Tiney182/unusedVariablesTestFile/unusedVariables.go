package main

import "fmt"

var gvar int // not an error ?

func main() {
	var one int
	two := 2
	var three int
	thee = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}
