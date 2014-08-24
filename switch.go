package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Run on OS ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Print("OSX")
	case "linux":
		fmt.Print("Linux")
	}
}
