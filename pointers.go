package main

import (
	"fmt"
)

func main() {
	i, j := 12, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}