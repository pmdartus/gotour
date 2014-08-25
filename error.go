package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	Why  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.Why)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't worked",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
