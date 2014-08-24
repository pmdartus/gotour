package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Simple for loop", sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("While loop", sum)

}
