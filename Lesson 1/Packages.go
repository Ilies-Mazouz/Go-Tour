// every GO program is made up of packages
// program runs in package main
package main


// this program is using the packages with import paths "fmt" and "math/rand"
// Convention: the package name is the same as the last element of the import path
// For instance: the "math/rand" package comprises files that begin with the statement package rand
import (
	"fmt"
	"math/rand"
)

// func main is the entry point for every GO program
//keyword func: defines a function, in javascript this would be function name() {}
func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Random number:", rand.Intn(10))
}

// The "fmt" package is used for formatting and printing output to the console.
// The "math/rand" package provides functions for generating random numbers.