package argparser

import (
	"fmt"
	"hamgo/lib/logger"
)

type Handler func(args []string)

type Command struct {
	handler     Handler
	description string
	usage       string
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		"help": {
			PrintHelp,
			"Show usage and help",
			"hamgo help",
		},
		"start": {
			logger.BeginSession,
			"Prints passed args",
			"hamgo start <PATH>",
		},
	}
}

func ParseArgs(args []string) {
	if len(args) == 1 {
		PrintHelp(make([]string, 0))
		return
	}

	first := args[1]

	command, ok := Commands[first]

	if !ok {
		fmt.Println("Invalid command")
		PrintHelp(make([]string, 0))
		return
	}

	command.handler(args[2:])
}
