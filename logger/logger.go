package logger

import (
	"bufio"
	"hamgo/types"
	"hamgo/util"
	"os"
)

func BeginSession(args []string) {
	if len(args) == 0 {
		println("Please specify a log file")
		return
	}

	path := args[0]

	mainLoop(openOrCreateFile(path))
}

func mainLoop(file *os.File) {
	writer := bufio.NewWriter(file)

	for {
		qso := RecordInput()
		qso.Write(writer)
	}
}

func RecordInput() types.QSO {
	return types.QSO{
		Date:        util.Prompt("Date"),
		Time:        util.Prompt("Time"),
		Freq:        util.Prompt("Frequency"),
		Mode:        util.Prompt("Mode"),
		Call:        util.Prompt("Call"),
		RstSent:     util.Prompt("RST Sent"),
		RstReceived: util.Prompt("RST Received"),
		Comment:     util.Prompt("Comment"),
	}
}

func openOrCreateFile(path string) *os.File {
	if fp, err := os.Open(path); err == nil {
		return fp
	}

	if fp, err := os.Create(path); err == nil {
		util.WriteHeader(bufio.NewWriter(fp))
		return fp
	}

	panic("Failed to create a file")
}
