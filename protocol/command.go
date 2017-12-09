package protocol

import (
	"encoding/json"
	"sync"
)

/*
CommandIface is the interface definition for a socket command
*/
type CommandIface interface {
	// Done handles the command result
	Done(result []byte, err error)
}

/*
Command is a generic CommandIface type
*/
type Command struct {
	Err    error
	Method string
	Params interface{}
	Result interface{}
	WG     sync.WaitGroup
}

/*
CommandResult represents the result of the socket command
*/
type CommandResult struct {
	Data string `json:"data"`
}

/*
Done is a SocketCmdIface implementation
*/
func (cmd *Command) Done(result []byte, err error) {
	if err == nil {
		err = json.Unmarshal(result, &cmd.Result)
	}
	cmd.Err = err
	cmd.WG.Done()
}

/*
CommandPayload is a representation of a WebSocket JSON payload
*/
type CommandPayload struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

/*
NewCommandPayload generates a new CommandPayload pointer
*/
func NewCommandPayload(id int, method string, params interface{}) *CommandPayload {
	return &CommandPayload{
		ID:     id,
		Method: method,
		Params: params,
	}
}
