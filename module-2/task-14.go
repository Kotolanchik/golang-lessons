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

	substr, _ := reader.ReadString('\n')
	substr = strings.TrimSpace(substr)

	index := strings.Index(text, substr)

	if index != -1 {
		fmt.Print(index)
	} else {
		fmt.Print(-1)
	}
}
