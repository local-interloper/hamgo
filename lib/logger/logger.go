package logger

import (
	"bufio"
	"hamgo/lib/types"
	"hamgo/lib/util"
	"os"
	"time"
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
	now := time.Now()
	return types.QSO{
		Date:        now.Format("YYYYMMDD"),
		Time:        now.Format("HHmm"),
		Freq:        util.Prompt("Frequency"),
		Mode:        util.Prompt("Mode"),
		Call:        util.Prompt("Call"),
		RstSent:     util.Prompt("RST Sent"),
		RstReceived: util.Prompt("RST Received"),
		Comment:     util.Prompt("Comment"),
	}
}

func openOrCreateFile(path string) *os.File {
	if fp, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644); err == nil {
		return fp
	}

	if fp, err := os.Create(path); err == nil {
		util.WriteHeader(bufio.NewWriter(fp))
		return fp
	}

	panic("Failed to create a file")
}
