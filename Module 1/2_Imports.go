package main

// this program is using the packages with import paths "fmt" and "math/rand"
// Convention: the package name is the same as the last element of the import path
// For instance: the "math/rand" package comprises files that begin with the statement package rand
// This code groups the imports into a parenthesized, "factored" import statement
// You can also import packages one at a time like this:
// import "fmt"
// import "math"
// But it is good style to use the factored import statement.

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}