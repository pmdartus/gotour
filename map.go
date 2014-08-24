package main

import "fmt"

type Vertex struct {
	X, Y float64
}

var m = map[string]Vertex{
	"Bell labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
