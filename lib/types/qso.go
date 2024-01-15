package types

import (
	"bufio"
	"fmt"
	"hamgo/lib/util"
	"strings"
	"time"
)

type QSO struct {
	Operator    string
	Station     string
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
		util.TagFrom("OPERATOR", qso.Date),
		util.TagFrom("STATION_CALLSIGN", qso.Date),
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

func NewQSOFromPrompt(session *Session) QSO {
	now := time.Now()

	qso := QSO{
		Operator:    session.Operator,
		Station:     session.StationCallsign,
		Date:        fmt.Sprintf("%.4d%.2d%.2d", now.Year(), now.Month(), now.Day()),
		Time:        fmt.Sprintf("%.2d%.2d", now.Hour(), now.Minute()),
		Freq:        util.PromptDefault("Frequency", session.Frequency),
		Mode:        util.PromptDefault("Mode", session.Mode),
		Call:        util.Prompt("Call"),
		RstSent:     util.Prompt("RST Sent"),
		RstReceived: util.Prompt("RST Received"),
		Comment:     util.Prompt("Comment"),
	}

	session.Frequency = qso.Freq
	session.Mode = qso.Mode

	return qso
}
