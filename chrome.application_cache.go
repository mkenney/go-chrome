package chrome

import (
	"app/chrome/protocol"
	"encoding/json"
	applicationCache "app/chrome/application_cache"

	log "github.com/Sirupsen/logrus"
)

/*
ApplicationCache: https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/
EXPERIMENTAL
*/
type ApplicationCache struct{}

/*
Enable enables application cache domain notifications.
*/
func (ApplicationCache) Enable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "ApplicationCache.enable"
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables application cache domain notifications.
*/
func (ApplicationCache) Disable(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "ApplicationCache.disable"
	socket.SendCommand(command)
	return command.Err
}

/*
GetFramesWithManifests returns array of frame identifiers with manifest urls for each frame
containing a document associated with some application cache.
*/
func (ApplicationCache) GetFramesWithManifests(socket *Socket) error {
	command := new(protocol.Command)
	command.method = "ApplicationCache.getFramesWithManifests"
	socket.SendCommand(command)
	return command.Err
}

/*
GetManifestForFrame returns manifest URL for document in the given frame.
*/
func (ApplicationCache) GetManifestForFrame(socket *Socket, params *applicationCache.GetManifestForFrameParams) error {
	command := new(protocol.Command)
	command.method = "ApplicationCache.getManifestForFrame"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
GetApplicationCacheForFrame returns relevant application cache data for the document
in given frame.
*/
func (ApplicationCache) GetApplicationCacheForFrame(socket *Socket, params *applicationCache.GetApplicationCacheForFrameParams) error {
	command := new(protocol.Command)
	command.method = "ApplicationCache.getManifestForFrame"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
OnApplicationCacheStatusUpdated adds a handler to the ApplicationCache.applicationCacheStatusUpdated
event.
*/
func OnApplicationCacheStatusUpdated(socket *Socket, callback func(event *applicationCache.ApplicationCacheStatusUpdatedEvent)) error {
	handler := protocol.NewEventHandler(
		"ApplicationCache.applicationCacheStatusUpdated",
		func(name string, params []byte) {
			event := new(applicationCache.ApplicationCacheStatusUpdatedEvent)
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		}
	)
	socket.AddEventHandler(handler)
	return command.Err
}
