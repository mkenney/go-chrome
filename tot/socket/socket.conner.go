package socket

import (
	"fmt"

	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
	"github.com/mkenney/go-chrome/codes"
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

	log.WithFields(log.Fields{
		"socketID": socket.socketID,
		"url":      socket.url.String(),
	}).Debug("connecting")
	websocket, err := socket.newSocket(socket.url)
	if nil != err {
		log.WithFields(log.Fields{
			"error":    err.Error(),
			"socketID": socket.socketID,
		}).Debug("received error")
		socket.connected = false
		return errs.Wrap(err, codes.SocketEventHandlerNotFound, "Connect() failed while creating socket")
	}

	socket.conn = websocket
	socket.connected = true

	log.WithFields(log.Fields{
		"socketID": socket.socketID,
		"url":      socket.url.String(),
	}).Debug("connection established")
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
		err = errs.Wrap(err, codes.SocketCloseFailed, "could not close socket connection")
	}
	socket.conn = nil
	socket.connected = false
	return err
}

/*
ReadJSON reads data from a websocket connection.

ReadJSON is a Conner implementation.
*/
func (socket *Socket) ReadJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return errs.Wrap(err, codes.SocketNotConnected, "not connected")
	}

	err = socket.conn.ReadJSON(&v)
	if nil != err {
		return errs.Wrap(err, codes.SocketReadFailed, "socket read failed")
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
		return errs.Wrap(err, codes.SocketNotConnected, "not connected")
	}

	err = socket.conn.WriteJSON(v)
	if nil != err {
		return errs.Wrap(err, codes.SocketWriteFailed, "socket write failed")
	}

	return nil
}
