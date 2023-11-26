package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a = a * a
	fmt.Println(a)

	// Использование rune для представления Unicode-символа
	var myRune rune = 'A'

	fmt.Printf("Value: %c\n", myRune)

	// Использование цикла для работы с строкой как с рунами
	str := "Привет, мир!"
	for _, char := range str {
		fmt.Printf("%c ", char)
	}
	fmt.Println()
}

// var a,b, c int
// в функции c := 1
// var a = 1
