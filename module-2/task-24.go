package main

import "fmt"

func main() {
	fmt.Println(divide(15, 5))
	defer fmt.Println("Program has been finished")
	fmt.Println(dividee(4, 0)) // момент паники
}
func dividee(x, y float64) float64 {
	if y == 0 {
		panic("division by zero!")
	}
	return x / y
}

/*
Если defer находится выше panic (ф-ии, которая вызывает панику), то сначала
выполнится дефер, а потом паника. Однако, если defer находится ниже panic,
то defer не выполнится вовсе!
*/
//Оператор defer позволяет выполнить определенную операцию после каких-то действий (даже если сработает panic),
//при этом не важно, где в реальности вызывается эта функция.
