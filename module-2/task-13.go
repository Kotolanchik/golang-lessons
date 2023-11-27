package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	runes := []rune(text)
	var newStr string
	for i := len(runes) - 1; i >= 0; i-- {
		//fmt.Printf("%c ", runes[i])

		newStr += string(runes[i])
	}
	if newStr == text {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")

	}
}
