package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func Fail() error {
	return errors.New("Oh no...")
}
