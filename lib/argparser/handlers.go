package argparser

import "fmt"

func PrintHelp(args []string) {
	println("Usage: hamgo <command> [args]")
	fmt.Println("Commands:")
	for k, v := range Handlers {
		fmt.Println(k, "-", v.description)
	}
}

func PrintArgs(args []string) {
	for _, e := range args {
		fmt.Println(e)
	}
}
