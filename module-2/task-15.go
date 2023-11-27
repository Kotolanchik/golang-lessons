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
	for i := 0; i < len(runes); i++ {
		//fmt.Printf("%c ", runes[i])
		if i%2 == 1 {
			newStr += string(runes[i])
		}
	}

	fmt.Print(newStr)
}
