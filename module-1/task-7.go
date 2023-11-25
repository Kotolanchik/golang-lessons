package main

import "fmt"

func main() {
	var inp int
	fmt.Scan(&inp)

	if inp == 0 {
		fmt.Print("Ноль")
	} else if inp > 0 {
		fmt.Print("Число положительное")
	} else {
		fmt.Print("Число отрицательное")
	}
}
