package socket

import (
	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
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
	if nil != socket.conn {
		return nil
	}

	socket.logger.Debug("connecting")
	websocket, err := socket.newSocket(socket.url)
	if nil != err {
		socket.logger.WithFields(log.Fields{"error": err.Error()}).
			Debug("received error")
		socket.mux.Lock()
		socket.conn = nil
		socket.mux.Unlock()
		return errs.Wrap(err, 0, "Connect() failed while creating socket")
	}
	socket.mux.Lock()
	socket.conn = websocket
	socket.mux.Unlock()

	socket.logger.Debug("connection established")
	return nil
}

/*
Connected returns whether a connection exists.

Connected is a Conner implementation.
*/
func (socket *Socket) Connected() bool {
	connected := nil != socket.conn
	return connected
}

/*
Disconnect closes a websocket connection.

Disconnect is a Conner implementation.
*/
func (socket *Socket) Disconnect() error {
	var err error
	socket.mux.Lock()
	if nil == socket.conn {
		socket.mux.Unlock()
		return errs.New(0, "not connected")
	}
	err = socket.conn.Close()
	socket.mux.Unlock()

	if nil != err {
		err = errs.New(0, "could not close socket connection")
	}

	socket.mux.Lock()
	socket.conn = nil
	socket.mux.Unlock()

	return err
}

/*
ReadJSON reads data from a websocket connection.

ReadJSON is a Conner implementation.
*/
func (socket *Socket) ReadJSON(v interface{}) error {
	socket.mux.Lock()
	if nil == socket.conn {
		socket.mux.Unlock()
		return errs.New(0, "not connected")
	}
	socket.mux.Unlock()

	err := socket.conn.ReadJSON(&v)

	if nil != err {
		return errs.Wrap(err, 0, "socket read failed")
	}

	return nil
}

/*
WriteJSON writes data to a websocket connection.

WriteJSON is a Conner implementation.
*/
func (socket *Socket) WriteJSON(v interface{}) error {
	socket.mux.Lock()
	defer socket.mux.Unlock()

	if nil == socket.conn {
		return errs.New(0, "not connected")
	}

	err := socket.conn.WriteJSON(v)
	if nil != err {
		return errs.Wrap(err, 0, "socket write failed")
	}

	return nil
}
