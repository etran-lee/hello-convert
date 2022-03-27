package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//subcommands
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	//flags
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
		fmt.Println("listed")
		os.Exit(1)
	}

	if flag.Parsed() {
		switch lang {
		default:
			fmt.Printf("%s", languages.English)
		}
	}
	flag.PrintDefaults()
}
