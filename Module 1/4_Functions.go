package main

import "fmt"

// func is pretty much the same as in js, it takes 0 or more arguments
// unlike most other languages, GO requires you to specify the type. 
// This always happens after the variable/function name
func add(x int, y int) int {
	return x + y
}

// another way of writing the add function is like this:
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}