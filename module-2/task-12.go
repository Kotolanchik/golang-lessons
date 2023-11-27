package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	runes := []rune(text)
	if isValidString(runes) {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}

func isValidString(s []rune) bool {
	return len(s) > 0 && unicode.IsUpper(s[0]) && strings.HasSuffix(string(s), ".")
}
