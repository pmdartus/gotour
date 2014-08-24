package main

import "fmt"
import "math"

func sqrt(i float64) string {
	if i < 0 {
		return sqrt(-i) + "i"
	}
	return fmt.Sprint(math.Sqrt(i))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(2, 3, 10),
		pow(3, 3, 10),
	)
}
