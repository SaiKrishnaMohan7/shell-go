package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var built_in = map[string]bool{}

func main() {
	fmt.Print("$ ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input %w", err)
	}
	input = strings.TrimSpace(input)

	if built_in[input] {
		// handle
		return
	}
	fmt.Printf("%s: command not found", input)
}
