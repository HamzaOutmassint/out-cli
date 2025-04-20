package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var yellow = color.New(color.FgYellow).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()

func main() {
	fmt.Println()
	color.Green("  >_ OUT-CLI, For more information type `help`")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		currentDir, err := os.Getwd()
		if err != nil {
			color.Red("Error getting current directory:", err)
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
			fmt.Println(color.RedString("  >_  "), yellow(args[0]), color.RedString(" is not an out-cli command. See "), green("`help`"))
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
			fmt.Println(color.RedString("  >_  "), yellow(args[0]), color.RedString(" is not an out-cli command. See "), green("`help`"))
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
			fmt.Println(color.RedString("  >_  "), yellow(args[0]), color.RedString(" is not an out-cli command. See "), green("`help`"))
			fmt.Println()
		}
	}
}

// displayHelpMenu prints the help message for the CLI.
func displayHelpMenu() {
	fmt.Println()
	color.Blue("  Available commands:")

	fmt.Printf("    %s       		 - Show this help message (usage: `%s`)\n", yellow("help"), yellow("help"))

	fmt.Printf("    %s                  - Create one or more directories or files (usage: `%s`)\n", yellow("mkd"), yellow("mkd <path> [<path>...]"))

	fmt.Printf("    %s         		 - Change the current directory (usage: `%s`)\n", yellow("cd"), yellow("cd <path>"))

	fmt.Printf("    %s         		 - List files in a directory (usage: `%s`)\n", yellow("ls"), yellow("ls <path>"))

	fmt.Printf("    %s         		 - Delete one or more files (usage: `%s`)\n", yellow("rm"), yellow("rm <path> [<path>...]"))

	fmt.Printf("    %s                  - Delete one or more directories (usage: `%s`)\n", yellow("rmd"), yellow("rmd <path> [<path>...]"))

	fmt.Printf("    %s          		 - Quickly navigate to the home directory (usage: `%s`)\n", yellow("/"), yellow("/"))

	fmt.Printf("    %s          	 - Clear the terminal screen (usage: `%s`)\n", yellow("cls"), yellow("cls"))

	fmt.Printf("    %s or %s       - Exit the program (usage: `%s` or `%s`)\n", yellow("exit"), yellow("Crtl+C"), yellow("exit"), yellow("Crtl+C"))

	fmt.Println()
}

// createDirectory creates a new directory with the specified name.
func createDirectory(n []string) {
	for _, v := range n {
		err := os.Mkdir(v, 0744)

		if err != nil {
			fmt.Println()
			fmt.Printf(color.RedString("  >_  Error - mkd "), "`%s`", color.RedString(": Cannot create a file when that file already exists.\n"), yellow(v))
			fmt.Println()
		}
	}
}

// changeCurrentDirectory changes the current working directory to the specified path.
func changeCurrentDirectory(path string) {
	err := os.Chdir(path)

	if err != nil {
		fmt.Println()
		fmt.Println(color.RedString("  >_   Error - cd: Unable to access "), yellow(path), color.RedString(" - Directory not found.\n"))
		fmt.Println()
	}
}

// listDirectoryContents lists the files and directories in the specified path.
func listDirectoryContents(path string) {
	entries, err := os.ReadDir(path)

	fmt.Println()
	if err != nil {
		fmt.Println(color.RedString("  >_   Error - ls: Unable to access "), yellow(path), color.RedString(" - Directory not found.\n"))
	} else if len(entries) == 0 {
		color.Red("  >_  Empty directory")
	} else {
		for _, entry := range entries {
			info, err := entry.Info()

			if err != nil {
				fmt.Println(err)
			}

			if entry.IsDir() {
				fmt.Println("  >   ", info.ModTime().Format("02/01/2006"), green("               <dir>                  "), yellow(entry.Name()))
			} else {
				fmt.Println("  >   ", info.ModTime().Format("02/01/2006"), blue("               <file>                  "), yellow(entry.Name()))
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
			color.Red("  >_  Error: ", err)
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
			color.Red("  >_  Error: ", err)
			fmt.Println()
		}
	}
}

// changeToHomeDirectory changes the current directory to the user's home directory.
func changeToHomeDirectory() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		color.Red("  >_  Error: ", err)
		return
	}

	err = os.Chdir(homeDir)

	if err != nil {
		color.Red("  >_  Error: ", err)
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
