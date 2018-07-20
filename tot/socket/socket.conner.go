package socket

import (
	"fmt"

	errs "github.com/mkenney/go-errors"
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
		return errs.Wrap(err, "Connect() failed while creating socket")
	}

	socket.conn = websocket
	socket.connected = true

	log.Debugf("socket #%d - socket.Connect(): connection to %s established", socket.socketID, socket.url.String())
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
	if nil != err {
		socket.listenErr.With(err)
	}
	socket.conn = nil
	socket.connected = false
	if 0 == len(socket.listenErr) {
		return nil
	}
	return socket.listenErr
}

/*
ReadJSON reads data from a websocket connection.

ReadJSON is a Conner implementation.
*/
func (socket *Socket) ReadJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return errs.Wrap(err, "not connected")
	}

	err = socket.conn.ReadJSON(&v)
	if nil != err {
		return errs.Wrap(err, "socket read failed")
	}

	return nil
}

/*
WriteJSON writes data to a websocket connection.

WriteJSON is a Conner implementation.
*/
func (socket *Socket) WriteJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return errs.Wrap(err, "not connected")
	}

	err = socket.conn.WriteJSON(v)
	if nil != err {
		return errs.Wrap(err, "socket write failed")
	}

	return nil
}
