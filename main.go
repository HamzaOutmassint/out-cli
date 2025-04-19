package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		case "cls":
			clearScreen()
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
		case "rm":
			removeFile(args[1:])
		case "rmd":
			removeDirectory(args[1:])
		case "mkd":
			createDirectory(args[1:])
		default:
			fmt.Println()
			fmt.Printf("  >_  '%s' is not an out-cli command. See 'help'.\n", args[0])
			fmt.Println()
		}
	} else if len(args) > 1 {
		switch command {
		case "mkd":
			createDirectory(args[1:])
		case "rm":
			removeFile(args[1:])
		case "rmd":
			removeDirectory(args[1:])
		default:
			fmt.Println()
			fmt.Printf("  >_  '%s' is not an out-cli command. See 'help'.\n", args[0])
			fmt.Println()
		}
	}
}

// displayHelpMenu prints the help message for the CLI.
func displayHelpMenu() {
	fmt.Println()
	fmt.Println("  Available commands:")
	fmt.Println("    help       		  	  - Show this help message")
	fmt.Println("    mkd         		  - Create a directory/file or multi directories/files (usage: mkd <path> <path>)")
	fmt.Println("    cd         		  	  - Change the current directory (usage: cd <path>)")
	fmt.Println("    ls         		  	  - List files in a directory (usage: ls <path>)")
	fmt.Println("    rm         		  	  - Delete a file/files or an empty directory/directories (usage: rm <path> <path>)")
	fmt.Println("    rmd         		  - Delete a directory or multi directories (usage: rm <path> <path>)")
	fmt.Println("    /          		  	  - Change to home directory (usage: / )")
	fmt.Println("    cls          		  	  - clear terminal")
	fmt.Println("    exit or Ctrl+C        	  - Exit the program")
	fmt.Println()
}

// createDirectory creates a new directory with the specified name.
func createDirectory(n []string) {
	for _, v := range n {
		err := os.Mkdir(v, 0744)

		if err != nil {
			fmt.Println()
			fmt.Printf("  >_  Error - mkd '%s': Cannot create a file when that file already exists.\n", v)
			fmt.Println()
		}
	}
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

// listDirectoryContents lists the files and directories in the specified path.
func listDirectoryContents(path string) {
	entries, err := os.ReadDir(path)

	fmt.Println()
	if err != nil {
		fmt.Printf("  >_   Error - ls: Unable to access '%s' - Directory not found.\n", path)
	} else if len(entries) == 0 {
		fmt.Println("  >_  Empty directory")
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

// remove a file with a specified name.
func removeFile(n []string) {
	for _, v := range n {
		err := os.Remove(v)

		if err != nil {
			fmt.Println()
			fmt.Println("  >_  Error: ", err)
			fmt.Println()
		}
	}
}

// remove a directory with a specified name.
func removeDirectory(n []string) {
	for _, v := range n {
		err := os.RemoveAll(v)

		if err != nil {
			fmt.Println()
			fmt.Println("  >_  Error: ", err)
			fmt.Println()
		}
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

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
