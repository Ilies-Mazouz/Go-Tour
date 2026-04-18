package main

import (
	"fmt"
	"math"
)

// in Go, a name is exported if it begins with a capital letter
// Example: Pizza is an exported name, as is Pi, which is exported from the math package

// When importing a package, you can refer to its exported names
// Any "unexported" names are not accessible from outside the package

func main() {
	fmt.Println(math.Pi)
}

// "pi" doesn't work. "Pi" does, because it is exported.