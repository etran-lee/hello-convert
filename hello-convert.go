package main

import (
	"flag"
	"fmt"
	"os"

	languages "github.com/etran-lee/hello-convert/data_structs"
)

func main() {
	// list sub command + flags
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listPage := listCmd.Int("page", 1, "list page [1-7]")
	listAll := listCmd.Bool("all", false, "list all available languages")

	// default flags
	lang := flag.String("lang", "", "choose a language to say 'hello' in; you can list available languages by using the 'list' subcommand (example: 'hello-convert list')")

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "list":
			listCmd.Parse(os.Args[2:])
		default:
			flag.Parse()
		}
	}

	if listCmd.Parsed() {
		if *listAll {
			languages.ListAllLanguages()
			os.Exit(1)
		}
		languages.ListLanguages(*listPage)
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
