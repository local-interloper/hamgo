package types

import (
	"bufio"
	"hamgo/lib/util"
	"strings"
)

type QSO struct {
	Date        string
	Time        string
	Freq        string
	Mode        string
	Call        string
	RstSent     string
	RstReceived string
	Comment     string
}

func (qso QSO) Write(writer *bufio.Writer) {
	record := []string{
		util.TagFrom("QSO_DATE", qso.Date),
		util.TagFrom("TIME_ON", qso.Time),
		util.TagFrom("FREQ", qso.Freq),
		util.TagFrom("MODE", qso.Mode),
		util.TagFrom("CALL", qso.Call),
		util.TagFrom("RST_SENT", qso.RstSent),
		util.TagFrom("RST_RCVD", qso.RstReceived),
		util.TagFrom("COMMENT", qso.Comment),
		"<EOR>\n",
	}

	if _, err := writer.WriteString(strings.Join(record, "\n")); err != nil {
		panic("Failed to record the QSO")
	}

	if err := writer.Flush(); err != nil {
		panic("Failed to flush file buffer")
	}
}
