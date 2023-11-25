package main

import "fmt"

func main() {
	var inp int
	fmt.Scan(&inp)

	if inp > 0 && inp < 10 {
		fmt.Println(inp)
	} else if inp > 9 && inp < 100 {
		fmt.Println(inp / 10)
	} else if inp > 99 && inp < 1000 {
		fmt.Println(inp / 100)
	} else if inp > 999 && inp < 10000 {
		fmt.Println(inp / 1000)
	} else if inp > 9999 && inp < 100000 {
		fmt.Println(inp / 10000)
	}
}
