package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello world from gogs")
}

func Fail() error {
	return errors.New("Oh no...")
}
