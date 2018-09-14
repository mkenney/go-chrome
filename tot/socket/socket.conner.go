package socket

import (
	errs "github.com/bdlm/errors"
	"github.com/bdlm/log"
)

/*
Conn returns the current web socket pointer. Conn is a Conner
implementation.
*/
func (socket *Socket) Conn() WebSocketer {
	socket.Connect()
	return socket.conn
}

/*
Connect establishes a websocket connection. Connect is a Conner
implementation.
*/
func (socket *Socket) Connect() error {
	socket.mux.Lock()
	if nil != socket.conn {
		socket.mux.Unlock()
		return nil
	}
	socket.mux.Lock()

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

	// launch the connection circuit breaker
	go func() {
		<-socket.disconnect
		success := true
		socket.mux.Lock()
		if nil == socket.conn {
			socket.mux.Unlock()
			socket.logger.Warn("disconnect requested on nil connection")
		} else {
			socket.mux.Unlock()
			err := socket.conn.Close()
			if nil != err {
				success = false
			}
			socket.mux.Lock()
			socket.conn = nil
			socket.mux.Unlock()
		}
		socket.disconnect <- success

	}()
	return nil
}

/*
Connected returns whether a connection exists. Connected is a Conner
implementation.
*/
func (socket *Socket) Connected() bool {
	socket.mux.Lock()
	connected := nil != socket.conn
	socket.mux.Unlock()
	return connected
}

/*
Disconnect closes a websocket connection. Disconnect is a Conner
implementation.
*/
func (socket *Socket) Disconnect() error {
	socket.disconnect <- true
	if !<-socket.disconnect {
		return errs.New(0, "error disconnecting socket connection")
	}
	return nil
}

/*
ReadJSON reads data from a websocket connection. ReadJSON is a Conner
implementation.
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
WriteJSON writes data to a websocket connection. WriteJSON is a Conner
implementation.
*/
func (socket *Socket) WriteJSON(v interface{}) error {
	socket.mux.Lock()
	if nil == socket.conn {
		socket.mux.Unlock()
		return errs.New(0, "not connected")
	}
	socket.mux.Unlock()

	err := socket.conn.WriteJSON(v)
	if nil != err {
		return errs.Wrap(err, 0, "socket write failed")
	}

	return nil
}
