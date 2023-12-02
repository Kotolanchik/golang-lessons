package main

import (
	"errors"
	"fmt"
)

func divide_m(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func main() {
	numerator := 10.0
	denominator := 2.0

	result, err := divide_m(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%.2f / %.2f = %.2f\n", numerator, denominator, result)
	}
}
