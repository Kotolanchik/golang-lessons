package main

import "fmt"

func main() {
	var inp int
	fmt.Scan(&inp)

	if inp < 1000 && inp > 9999 {
		fmt.Println("NO")
		return
	}

	if (inp%4 == 0 && inp%100 != 0) || inp%400 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
