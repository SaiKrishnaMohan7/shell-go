package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var built_in = map[string]func(args string){
	"exit": handleExit,
}

const PROMPT = "$ "

func main() {
	for {
		fmt.Printf(PROMPT)

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input %w", err)
		}
		cmd := strings.TrimSpace(input)

		cmdHandler, ok := built_in[cmd]

		if ok {
			cmdHandler("")
			return
		}
		fmt.Printf("%s: command not found \n", cmd)
	}
}

func handleExit(args string) {
	os.Exit(0)
}
