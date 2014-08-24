package main

import (
	"fmt"
	"math"
)

func main() {
	// Explicit conversion
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z = int(f)
	fmt.Println(f, z)

	// Implicit conversions
	i := 3
	j := float64(i)
	k := uint(j)
	var format string = "%T - %v \n"
	fmt.Printf(format, i, i)
	fmt.Printf(format, j, j)
	fmt.Printf(format, k, k)
}
