package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("my error") // or fmt.Errorf("my error")
	fmt.Println("", err)
}
