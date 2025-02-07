package main

import (
	"fmt"
	"os"

	"github.com/dimmerz92/go-shadcn/internal"
)

const version = "v0.0.0"

const usage = `
	goshadcn <command> [<args>...]

	go-shadcn - a port of the shadcn components for Go

	commands:
	dev -p <port>
	-	Runs the dev server for component testing and development.
	
	help, -h, --h, -help, --help,
	<command> -h, <command> --h, <command> -help, <command> --help
	-	Prints usage information

	version, -v, --v, -version, --version
	-	Prints current version
	`

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	switch args[1] {

	case "dev":
		internal.DevServer(args[2:])

	case "help", "-h", "--h", "-help", "--help":
		fmt.Print(usage)

	case "version", "-v", "--v", "-version", "--version":
		fmt.Printf(version)

	default:
		fmt.Println(usage)
		os.Exit(1)
	}
}
