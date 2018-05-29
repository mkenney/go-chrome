package socket

import (
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

/*
Conn returns the current web socket pointer.

Conn is a Conner implementation.
*/
func (socket *Socket) Conn() WebSocketer {
	socket.Connect()
	return socket.conn
}

/*
Connect establishes a websocket connection.

Connect is a Conner implementation.
*/
func (socket *Socket) Connect() error {
	socket.mux.Lock()
	defer socket.mux.Unlock()

	if socket.connected {
		return nil
	}

	log.Debugf("socket #%d - socket.Connect(): connecting to %s", socket.socketID, socket.url.String())
	websocket, err := socket.newSocket(socket.url)
	if nil != err {
		log.Debugf("socket #%d - socket.Connect(): received error %s", socket.socketID, err.Error())
		socket.connected = false
		return errors.Wrap(err, "creating socket failed")
	}
	socket.conn = websocket

	log.Debugf("socket #%d - socket.Connect(): connection to %s established", socket.socketID, socket.url.String())
	socket.connected = true

	return nil
}

/*
Connected returns whether a connection exists.

Connected is a Conner implementation.
*/
func (socket *Socket) Connected() bool {
	return socket.connected
}

/*
Disconnect closes a websocket connection.

Disconnect is a Conner implementation.
*/
func (socket *Socket) Disconnect() error {
	if !socket.connected {
		return fmt.Errorf("not connected")
	}
	socket.Stop()
	err := socket.conn.Close()
	socket.conn = nil
	socket.connected = false
	return errors.Wrap(err, "disconnect failed")
}

/*
ReadJSON reads data from a websocket connection.

ReadJSON is a Conner implementation.
*/
func (socket *Socket) ReadJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return errors.Wrap(err, "socket connect failed")
	}
	return socket.conn.ReadJSON(&v)
}

/*
WriteJSON writes data to a websocket connection.

WriteJSON is a Conner implementation.
*/
func (socket *Socket) WriteJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return errors.Wrap(err, "socket connect failed")
	}
	return socket.conn.WriteJSON(v)
}
