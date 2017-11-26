package chrome

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type LoadEvent struct {
	Timestamp float64 `json:"timestamp"`
}

func OnLoadEvent(socket *Socket, fn func(evt *LoadEvent)) {
	socket.AddEventHandler("Page.loadEventFired", func(name string, params []byte) {
		evt := &LoadEvent{}
		if err := json.Unmarshal(params, evt); err != nil {
			log.Error(err)
		} else {
			fn(evt)
		}
	})
}
