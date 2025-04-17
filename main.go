package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Println("  >_ OUT-CLI, For more information type 'help'")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			break
		}

		fmt.Print(currentDir, " >_ ")

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

// execution executes the command based on the input string.
func execution(s string) {
	args := strings.Fields(s)
	command := args[0]

	if len(args) == 1 {
		switch command {
		case "help":
			printHelp()
		case "ls":
			listFiles(".")
		case "/":
			homeDire()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println()
			fmt.Printf("  >_  '%s' is not an out-cli command. See 'help'.\n", args[0])
			fmt.Println()
		}
	} else if len(args) > 1 && len(args) < 3 {
		switch command {
		case "ls":
			listFiles(args[1])
		case "cd":
			changeDirectory(args[1])
		}
	} else {
		fmt.Println("many args")
	}
}

// printHelp prints the help message for the CLI.
func printHelp() {
	fmt.Println()
	fmt.Println("  Available commands:")
	fmt.Println("    help       		  - Show this help message")
	fmt.Println("    cr         		  - Create a directory (usage: cr <directory>)")
	fmt.Println("    cd         		  - Changes the current directory (usage: cd <directory>)")
	fmt.Println("    ls         		  - List files in a directory (usage: ls)")
	fmt.Println("    rm         		  - Delete a file or directory (usage: rm <path>)")
	fmt.Println("    /          		  - Change to home directory (usage: / )")
	fmt.Println("    exit or Ctrl+C        - Exit the program")
	fmt.Println()
}

// listFiles lists the files in the specified directory.
func listFiles(n string) {
	enteries, err := os.ReadDir(n)

	fmt.Println()
	if err != nil {
		fmt.Printf("  >_   Error - ls: Unable to access '%s' - Directory not found.\n", n)
	} else {
		for _, entry := range enteries {
			info, err := entry.Info()

			if err != nil {
				fmt.Println(err)
			}

			if entry.IsDir() {
				fmt.Printf("  >   %s               <dir>                  %s\n", info.ModTime().Format("02/01/2006"), entry.Name())
			} else {
				fmt.Printf("  >   %s               <file>                 %s\n", info.ModTime().Format("02/01/2006"), entry.Name())
			}
		}
	}
	fmt.Println()
}

// changeDirectory changes the current directory to the specified path.
func changeDirectory(n string) {
	err := os.Chdir(n)

	if err != nil {
		fmt.Println()
		fmt.Printf("  >_   Error - cd: Unable to access '%s' - Directory not found.\n", n)
		fmt.Println()
	}
}

// homeDire changes the current directory to the user's home directory.
func homeDire() {
	targetPath, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = os.Chdir(targetPath)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
