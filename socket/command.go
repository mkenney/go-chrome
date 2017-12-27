package socket

import (
	"encoding/json"
	"errors"
	"sync"

	log "github.com/sirupsen/logrus"
)

/*
Command represents a command delivered to socket
*/
type Command struct {
	// Err contains any error resulting from executing the command
	Err error

	// Method is the Chrome protocol method being executed
	Method string

	// Optional. Params holds the parameter struct for the command being executed
	Params interface{}

	// Optional. Result holds the result struct for the command being executed
	Result interface{}

	// WG handles command syncronization
	WG sync.WaitGroup
}

/*
Done manages the result of a command and signals completion
*/
func (cmd *Command) Done(result []byte, err error) {
	if err == nil {
		err = json.Unmarshal(result, &cmd.Result)
	}
	cmd.Err = err
	cmd.WG.Done()
}

/*
CommandPayload represents a WebSocket JSON payload for a command
*/
type CommandPayload struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

/*
SendCommand sends a command to a connected socket.

The socket's command mutex is locked, the command counter is incremented, the payload is sent to the
socket connection and the mutex is unlocked. The command is stored using the counter value as it's
Id. When the command is executed and the socket responds, handleCmd() is executed to generate a
response.
*/
func (socket *Socket) SendCommand(command *Command) int {

	// Safely get generate a command ID
	socket.Mux.Lock()
	socket.CommandID++
	commandID := socket.CommandID
	socket.Mux.Unlock()

	// Safely add a command to the internal stack
	socket.Commands.Mux.Lock()
	payload := &CommandPayload{
		ID:     commandID,
		Method: command.Method,
		Params: command.Params,
	}
	socket.Commands.Map[payload.ID] = command
	socket.Commands.Mux.Unlock() // Do not defer, handlCommand also locks

	tmp, _ := json.Marshal(payload)
	log.Debugf("Sending command '%s'", string(tmp))

	command.WG.Add(1)
	if err := socket.conn.WriteJSON(payload); err != nil {
		command.Done(nil, err)
	}
	command.WG.Wait()

	return payload.ID
}

func (socket *Socket) handleCommand(response *SocketResponse) {
	var err error

	socket.Commands.Mux.Lock()
	defer socket.Commands.Mux.Unlock()

	if command, ok := socket.Commands.Map[response.ID]; !ok {
		log.Errorf("Command %d not found: result=%s err=%s", response.ID, response.Result, response.Error.Message)

	} else {
		delete(socket.Commands.Map, response.ID)
		if "" != response.Error.Message {
			err = errors.New(response.Error.Message)
		}
		command.Done(response.Result, err)
	}
}
