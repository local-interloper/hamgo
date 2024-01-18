package argparser

import "fmt"

func PrintHelp(args []string) {
	println("Usage: hamgo <command> [args]")
	fmt.Println("Commands:")
	for k, v := range Commands {
		fmt.Println(k, "-", v.description)
	}
}
