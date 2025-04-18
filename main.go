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

		executeCommand(input)
	}
}

// executeCommand parses and executes the user's input command.
func executeCommand(input string) {
	args := strings.Fields(input)
	command := args[0]

	if len(args) == 1 {
		switch command {
		case "help":
			displayHelpMenu()
		case "ls":
			listDirectoryContents(".")
		case "/":
			changeToHomeDirectory()
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
			listDirectoryContents(args[1])
		case "cd":
			changeCurrentDirectory(args[1])
		}
	} else {
		fmt.Println("Too many arguments.")
	}
}

// displayHelpMenu prints the help message for the CLI.
func displayHelpMenu() {
	fmt.Println()
	fmt.Println("  Available commands:")
	fmt.Println("    help       		  - Show this help message")
	fmt.Println("    cr         		  - Create a directory (usage: cr <directory>)")
	fmt.Println("    cd         		  - Change the current directory (usage: cd <directory>)")
	fmt.Println("    ls         		  - List files in a directory (usage: ls [directory])")
	fmt.Println("    rm         		  - Delete a file or directory (usage: rm <path>)")
	fmt.Println("    /          		  - Change to home directory (usage: / )")
	fmt.Println("    exit or Ctrl+C        - Exit the program")
	fmt.Println()
}

// listDirectoryContents lists the files and directories in the specified path.
func listDirectoryContents(path string) {
	entries, err := os.ReadDir(path)

	fmt.Println()
	if err != nil {
		fmt.Printf("  >_   Error - ls: Unable to access '%s' - Directory not found.\n", path)
	} else {
		for _, entry := range entries {
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

// changeCurrentDirectory changes the current working directory to the specified path.
func changeCurrentDirectory(path string) {
	err := os.Chdir(path)

	if err != nil {
		fmt.Println()
		fmt.Printf("  >_   Error - cd: Unable to access '%s' - Directory not found.\n", path)
		fmt.Println()
	}
}

// changeToHomeDirectory changes the current directory to the user's home directory.
func changeToHomeDirectory() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = os.Chdir(homeDir)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
