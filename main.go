package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("OUT-CLI, For more information type 'help'")

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

		execution(input)
	}
}

func execution(s string){
	args := strings.Fields(s)
	command := args[0]

	if len(args) == 1 {
		switch command {
			case "help":
				printHelp()
			case "ls":
				listFiles(".")
			default: 
				fmt.Printf("'%s' is not an out-cli command. See 'help'.\n",args[0])
		}
	}else if len(args) > 1 && len(args) < 3 {
		fmt.Println(command)
		fmt.Println(args[1])

		switch command{
			case "ls":
				listFiles(args[1])
		}
	}else {
		fmt.Println("many args")
	}

}

func printHelp(){
	fmt.Println()
	fmt.Println("Available commands:")
    fmt.Println("  help       		- Show this help message")
    fmt.Println("  cr         		- Create a directory (usage: cr <directory>)")
    fmt.Println("  cd         		- changes the current directory (usage: cd <directory>)")
    fmt.Println("  ls         		- List files in a directory (usage: ls)")
    fmt.Println("  rm         		- Delete a file or directory (usage: rm <path>)")
    fmt.Println("  exit or Ctrl+C        - Exit the program")
	fmt.Println()
}

func listFiles(s string){

}