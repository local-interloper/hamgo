package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var streamReader *bufio.Reader

func init() {
	streamReader = bufio.NewReader(os.Stdin)
}

func WriteHeader(writer *bufio.Writer) {
	text := []string{
		TagFrom("ADIF_VER", "3.0.6"),
		TagFrom("PROGRAMID", "hamgo"),
		"<EOH>\n",
	}

	if _, err := writer.WriteString(strings.Join(text, "\n")); err != nil {
		panic("Failed to write the file header")
	}

	if err := writer.Flush(); err != nil {
		panic("Failed to flush the header to log")
	}
}

func GetLine() string {
	for {
		text, err := streamReader.ReadString('\n')

		if err != nil {
			fmt.Println("Could not read standard input, please try again.")
			continue
		}

		return strings.TrimSpace(text)
	}
}

func Prompt(prompt string) string {
	fmt.Print(prompt + ": ")
	return GetLine()
}

func PromptDefault(prompt string, def string) string {
	input := Prompt(fmt.Sprintf("%s [%s]", prompt, def))

	if len(input) == 0 {
		input = def
	}

	return input
}
