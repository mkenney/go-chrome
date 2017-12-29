package protocol

import (
	"encoding/json"

	applicationCache "github.com/mkenney/go-chrome/protocol/application_cache"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
ApplicationCache provides a namespace for the Chrome Animation protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/ EXPERIMENTAL.
*/
var ApplicationCache = ApplicationCacheProtocol{}

/*
ApplicationCacheProtocol defines the ApplicationCache protocol methods.
*/
type ApplicationCacheProtocol struct{}

/*
Enable enables application cache domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-enable
*/
func (ApplicationCacheProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("ApplicationCache.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetForFrame returns relevant application cache data for the document
in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
func (ApplicationCacheProtocol) GetForFrame(
	socket sock.Socketer,
	params *applicationCache.GetForFrameParams,
) (*applicationCache.GetForFrameResult, error) {
	command := sock.NewCommand("ApplicationCache.getApplicationCacheForFrame", params)
	result := &applicationCache.GetForFrameResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetFramesWithManifests returns array of frame identifiers with manifest urls for each frame
containing a document associated with some application cache.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
*/
func (ApplicationCacheProtocol) GetFramesWithManifests(
	socket sock.Socketer,
) (*applicationCache.GetFramesWithManifestsResult, error) {
	command := sock.NewCommand("ApplicationCache.getFramesWithManifests", nil)
	result := &applicationCache.GetFramesWithManifestsResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetManifestForFrame returns manifest URL for document in the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
*/
func (ApplicationCacheProtocol) GetManifestForFrame(
	socket sock.Socketer,
	params *applicationCache.GetManifestForFrameParams,
) (*applicationCache.GetManifestForFrameResult, error) {
	command := sock.NewCommand("ApplicationCache.getManifestForFrame", params)
	result := &applicationCache.GetManifestForFrameResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
OnApplicationCacheStatusUpdated adds a handler to the ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
*/
func (ApplicationCacheProtocol) OnApplicationCacheStatusUpdated(
	socket sock.Socketer,
	callback func(event *applicationCache.StatusUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ApplicationCache.applicationCacheStatusUpdated",
		func(response *sock.Response) {
			event := &applicationCache.StatusUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnNetworkStateUpdated adds a handler to the ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
*/
func (ApplicationCacheProtocol) OnNetworkStateUpdated(
	socket sock.Socketer,
	callback func(event *applicationCache.StatusUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ApplicationCache.networkStateUpdated",
		func(response *sock.Response) {
			event := &applicationCache.StatusUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
