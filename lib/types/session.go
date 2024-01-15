package types

import (
	"hamgo/lib/util"
)

type Session struct {
	Operator        string
	StationCallsign string
	Frequency       string
	Mode            string
}

func NewSessionFromPrompt() Session {
	operator := util.Prompt("Operator callsign")
	station := util.PromptDefault("Station callsign", operator)

	if len(station) == 0 {
		station = operator
	}

	return Session{operator, station, "3.5", "LSB"}
}
