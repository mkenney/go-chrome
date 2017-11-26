package chrome

import (
	"app/chrome/protocol"
	"encoding/json"
	appcache "app/chrome/protocol/application_cache"

	log "github.com/Sirupsen/logrus"
)

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
func (ApplicationCache) GetManifestForFrame(socket *Socket, params *appcache.GetManifestForFrameParams) error {
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
func (ApplicationCache) GetApplicationCacheForFrame(socket *Socket, params *appcache.GetApplicationCacheForFrameParams) error {
	command := new(protocol.Command)
	command.method = "ApplicationCache.getManifestForFrame"
	command.params = params
	socket.SendCommand(command)
	return command.Err
}

/*
OnStatusUpdated
*/
func OnStatusUpdated(socket *Socket, callback func(event *appcache.StatusUpdatedEvent)) error {
	event := protocol.NewEvent(
		"ApplicationCache.applicationCacheStatusUpdated",
		func(name string, params []byte) {
			event := new(ApplicationCacheForFrameEvent)
			if err := json.Unmarshal(params, evt); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		}
	)
	socket.AddEventHandler(event)
	return command.Err
}
