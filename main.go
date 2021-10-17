package main

import (
	"fmt"
	"os"

	"github.com/jgrancell/kubenv/cmd"
	"github.com/jgrancell/kubenv/settings"
)

func main() {
	Run(os.Args[1:])
}

func Run(args []string) {
	version := "1.0.0"

	// Loading configuration
	conf, err := settings.LoadConfiguration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(args) >= 1 {
		var result bool
		command := args[0]

		switch command {
			case "delete":
				result, err = cmd.Delete(args[1:], conf)
			case "import":
				result, err = cmd.Import(args[1:], conf)
			case "list":
				result, err = cmd.List(conf)
			case "use":
				result, err = cmd.Activate(args[1:], conf)

			default:
				result, err = cmd.Help(version)
		}
		if result {
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		_, _ = cmd.Help(version)
		os.Exit(1)
	}
}
