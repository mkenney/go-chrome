package socket

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

/*
Conn returns a struct that implements the WebSocketer interface.
*/
func (socket *Socket) Conn() *ChromeWebSocket {
	socket.Connect()
	return socket.conn
}

/*
Connect implements Conner.
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
		return err
	}
	socket.conn = websocket

	log.Debugf("socket #%d - socket.Connect(): connection to %s established", socket.socketID, socket.url.String())
	socket.connected = true

	return nil
}

/*
Connected implements Conner.
*/
func (socket *Socket) Connected() bool {
	return socket.connected
}

/*
Disconnect implements Conner.
*/
func (socket *Socket) Disconnect() error {
	if !socket.connected {
		return fmt.Errorf("Could not disconnect (no connection exists)")
	}
	err := socket.conn.Close()
	socket.conn = nil
	socket.connected = false
	return err
}

/*
ReadJSON implements Conner.
*/
func (socket *Socket) ReadJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return err
	}
	return socket.conn.ReadJSON(&v)
}

/*
WriteJSON implements Conner.
*/
func (socket *Socket) WriteJSON(v interface{}) error {
	err := socket.Connect()
	if nil != err {
		return err
	}
	return socket.conn.WriteJSON(v)
}
