package main

import "fmt"

func main() {
	var inp int
	fmt.Scan(&inp)

	if inp < 100000 && inp > 999999 {
		fmt.Println("NO")
		return
	}

	first := inp / 100000
	second := (inp / 10000) % 10
	third := (inp / 1000) % 10
	fourth := (inp / 100) % 10
	fifth := (inp % 100) / 10
	sixth := inp % 10

	if first+second+third == fourth+fifth+sixth {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// fmt.Println(fmt.Sprintf("%d %d %d %d %d %d ", first, second, third, fourth, fifth, sixth))
}
