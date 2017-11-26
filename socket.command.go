package chrome

import (
	"app/chrome/protocol"
	"encoding/json"
	"errors"

	log "github.com/Sirupsen/logrus"
)

/*
SendCommand sends a command to a connected socket.

The socket's command mutex is locked, the command counter is incremented, the payload is sent to the
socket connection and the mutex is unlocked. The command is stored using the counter value as it's
Id. When the command is executed and the socket responds, handleCmd() is executed to generate a
response.
*/
func (socket *Socket) SendCommand(command protocol.CommandIface) int {
	command.WG.Add(1)
	socket.cmdMutex.Lock()
	defer socket.cmdMutex.Unlock()

	socket.cmdID++
	payload := &protocol.CommandPayload{
		socket.cmdID,
		command.Method(),
		command.Params(),
	}
	tmp, _ := json.Marshal(payload)
	log.Debugf("Sending %#v", string(tmp))
	if err := socket.conn.WriteJSON(payload); err != nil {
		command.Done(nil, err)
	}
	socket.commands[payload.ID] = command

	command.WG.Wait()
	return payload.ID
}

func (socket *Socket) handleCommand(response *SocketResponse) {
	var err error
	log.Infof("Command #%d responding", response.ID)
	socket.cmdMutex.Lock()
	defer socket.cmdMutex.Unlock()

	if command, ok := socket.commands[response.ID]; !ok {
		log.Warnf("Command %d not found: result=%s err=%s", response.ID, response.Result, response.Error.Message)

	} else {
		delete(socket.commands, response.ID)

		if "" != response.Error.Message {
			err = errors.New(response.Error.Message)
		}

		command.Done(response.Result, err)

		tmp, _ := json.Marshal(response)
		log.Debugf("Response: %s", tmp)
		tmp, _ = json.Marshal(command)
		log.Debugf("Command: %s", tmp)
	}
}
