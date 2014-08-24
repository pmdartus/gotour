package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("Answer is", m["Answer"])

	m["Answer"] = 48
	fmt.Println("Answer is", m["Answer"])

	delete(m, "Answer")

	v, ok := m["Answer"]
	fmt.Println(v, ok)
}
