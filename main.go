package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// CommandHandler defines the signature for command handler functions.
type CommandHandler func(args []string) error

// commandRegistry maps commands to their respective handler functions.
var commandRegistry = map[string]CommandHandler{
	"help": displayHelpMenu,
	"ls":   listDirectoryContents,
	"/":    changeToHomeDirectory,
	"cls":  clearScreen,
	"exit": exitCLI,
	"cd":   changeCurrentDirectory,
	"mkd":  createDirectory,
	"mkf":  createFile,
	"rmf":  removeFile,
	"rmd":  removeDirectory,
}

// Utility Functions for Colored Output
func red(s string) string    { return color.New(color.FgRed).SprintFunc()(s) }
func yellow(s string) string { return color.New(color.FgYellow).SprintFunc()(s) }
func green(s string) string  { return color.New(color.FgGreen).SprintFunc()(s) }
func blue(s string) string   { return color.New(color.FgBlue).SprintFunc()(s) }

func main() {
	fmt.Println()
	color.Green("  >_ OUT-CLI, For more information type `help`")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		currentDir, err := os.Getwd()
		if err != nil {
			printError(fmt.Sprintf("Error getting current directory: %v", err))
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
	if len(args) == 0 {
		return
	}

	command := args[0]
	handler, exists := commandRegistry[command]
	if !exists {
		printError(fmt.Sprintln(yellow(args[0]), color.RedString(" is not an out-cli command. See "), green("`help`")))
		return
	}

	// Execute the command with the remaining arguments
	err := handler(args[1:])
	if err != nil {
		printError(err.Error())
	}
}

// printError prints an error message in red color.
func printError(message string) {
	fmt.Println()
	fmt.Println(red("  >_  Error "), message)
	fmt.Println()
}

// displayHelpMenu prints the help message for the CLI.
func displayHelpMenu(_ []string) error {
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
	return nil
}

// createDirectory creates one or more directories.
func createDirectory(paths []string) error {
	for _, path := range paths {
		err := os.Mkdir(path, 0744)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}
	return nil
}

// createFile creates one or more files.
func createFile(paths []string) error {
	for _, path := range paths {
		file, err := os.OpenFile(path, os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
		defer file.Close()
	}
	return nil
}

// changeCurrentDirectory changes the current working directory.
func changeCurrentDirectory(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: cd <path>")
	}
	return os.Chdir(args[0])
}

// listDirectoryContents lists the contents of a directory.
func listDirectoryContents(args []string) error {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		errMessage := fmt.Sprintln(color.RedString("- ls: Unable to access "), yellow(dir), color.RedString(" - Directory not found"))

		return fmt.Errorf("%s", errMessage)
	}

	fmt.Println()
	if len(entries) == 0 {
		fmt.Println(red("  >_  Empty directory"))
	} else {
		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				return fmt.Errorf("%s", err)
			}

			if entry.IsDir() {
				fmt.Println("  >   ", info.ModTime().Format("02/01/2006"), green("               <dir>                  "), yellow(entry.Name()))
			} else {
				fmt.Println("  >   ", info.ModTime().Format("02/01/2006"), blue("               <file>                 "), yellow(entry.Name()))
			}
		}
	}
	fmt.Println()
	return nil
}

// removeFile deletes one or more files.
func removeFile(paths []string) error {
	for _, path := range paths {
		err := os.Remove(path)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}
	return nil
}

// removeDirectory deletes one or more directories.
func removeDirectory(paths []string) error {
	for _, path := range paths {
		err := os.RemoveAll(path)
		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}
	return nil
}

// changeToHomeDirectory navigates to the user's home directory.
func changeToHomeDirectory(_ []string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return os.Chdir(homeDir)
}

// clearScreen clears the terminal screen.
func clearScreen(_ []string) error {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// exitCLI exits the program.
func exitCLI(_ []string) error {
	os.Exit(0)
	return nil
}
