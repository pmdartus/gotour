package main

import "fmt"

func split(sum int) (x int, y int) {
	y = sum * 4
	x = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
