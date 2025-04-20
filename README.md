# Out-CLI

Out-CLI is a custom command-line application built with Go, born out of the need for a more personalized and intuitive way to manage directories and files on Windows. The motivation behind creating this tool was simple: I wanted to replace the default Windows commands with my own custom commands—ones that are easier for me to remember and use. Over time, this project has become a way for me to tailor the command-line experience to my preferences, allowing me to add new commands whenever I need them, without relying on or struggling to recall the official Windows commands.

## Features

- **Directory Navigation**:
  - `cd <path>`: Change the current directory.cls
  - `/`: Quickly navigate to the home directory.

- **Directory and File Management**:
  - `mkd <path> [<path>...]`: Create one or more directories or files.
  - `rm <path> [<path>...]`: Delete one or more files.
  - `rmd <path> [<path>...]`: Delete one or more directories.

- **Directory Listing**:
  - `ls [<path>]`: List the contents of a directory. Defaults to the current directory if no path is provided.

- **Terminal Management**:
  - `cls`: Clear the terminal screen.

- **Help and Exit**:
  - `help`: Display the help menu with a list of available commands.
  - `exit`: Exit the application.

## Installation

1. Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org/).
2. Clone this repository:
   ```sh
   git clone https://github.com/hamzaOutmassint/out-cli.git
