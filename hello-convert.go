package main

import (
	"flag"
	"fmt"
	"os"

	languages "github.com/etran-lee/hello-convert/data_structs"
)

func main() {
	// subcommands
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// flags
	lang := flag.String("lang", "", "choose a language to say 'hello', you can list available languages by using the 'list' subcommand (example: 'hello-convert list')")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "list":
			listCmd.Parse(os.Args[2:])
		default:
			flag.Parse()
		}
	}

	if listCmd.Parsed() {
		languages.ListLanguages()
		os.Exit(1)
	}

	if flag.Parsed() {
		if len(*lang) < 3 {
			fmt.Print("must input a language key longer than 3 characters")
			os.Exit(1)
		}

		fmt.Printf("%s", languages.KeyCheck(*lang).Output)
		os.Exit(1)
	}

	flag.PrintDefaults()
}
