package main

import (
	"hamgo/lib/argparser"
	"os"
)

func main() {
	argparser.Handle(os.Args)
}
