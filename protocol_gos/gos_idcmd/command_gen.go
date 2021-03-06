// Code generated by "genprotocol -ver=68269c8bfeecd6e461aa862d64007f60a1aeccd64229d523ae7b99a446255112 -basedir=. -prefix=gos -statstype=int"

package gos_idcmd

import "fmt"

type CommandID uint16 // use in packet header, DO NOT CHANGE
const (
	Invalid     CommandID = iota //
	Login                        //
	Heartbeat                    //
	MakeStage                    //
	EnterStage                   //
	ChatToStage                  //
	LeaveStage                   //

	CommandID_Count int = iota
)

var _CommandID2string = [CommandID_Count][2]string{
	Invalid:     {"Invalid", ""},
	Login:       {"Login", ""},
	Heartbeat:   {"Heartbeat", ""},
	MakeStage:   {"MakeStage", ""},
	EnterStage:  {"EnterStage", ""},
	ChatToStage: {"ChatToStage", ""},
	LeaveStage:  {"LeaveStage", ""},
}

func (e CommandID) String() string {
	if e >= 0 && e < CommandID(CommandID_Count) {
		return _CommandID2string[e][0]
	}
	return fmt.Sprintf("CommandID%d", uint16(e))
}

func (e CommandID) CommentString() string {
	if e >= 0 && e < CommandID(CommandID_Count) {
		return _CommandID2string[e][1]
	}
	return ""
}

var _string2CommandID = map[string]CommandID{
	"Invalid":     Invalid,
	"Login":       Login,
	"Heartbeat":   Heartbeat,
	"MakeStage":   MakeStage,
	"EnterStage":  EnterStage,
	"ChatToStage": ChatToStage,
	"LeaveStage":  LeaveStage,
}

func String2CommandID(s string) (CommandID, bool) {
	v, b := _string2CommandID[s]
	return v, b
}
