package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("welcom to OUT-CLI")

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("out-cli > ")

		// Exit if there's an error reading input
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "" {
			continue
		}

		fmt.Println(input)
	}
}