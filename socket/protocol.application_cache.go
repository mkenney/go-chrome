package socket

import (
	"encoding/json"

	applicationCache "github.com/mkenney/go-chrome/protocol/application_cache"
	log "github.com/sirupsen/logrus"
)

/*
ApplicationCacheProtocol provides a namespace for the Chrome Animation protocol
methods.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/ EXPERIMENTAL.
*/
type ApplicationCacheProtocol struct {
	Socket Socketer
}

/*
Enable enables application cache domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-enable
*/
func (protocol *ApplicationCacheProtocol) Enable() error {
	command := NewCommand("ApplicationCache.enable", nil)
	protocol.Socket.SendCommand(command)
	return command.Error()
}

/*
GetForFrame returns relevant application cache data for the document
in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
func (protocol *ApplicationCacheProtocol) GetForFrame(
	params *applicationCache.GetForFrameParams,
) (*applicationCache.GetForFrameResult, error) {
	command := NewCommand("ApplicationCache.getApplicationCacheForFrame", params)
	result := &applicationCache.GetForFrameResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetFramesWithManifests returns array of frame identifiers with manifest urls for
each frame containing a document associated with some application cache.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
*/
func (protocol *ApplicationCacheProtocol) GetFramesWithManifests() (*applicationCache.GetFramesWithManifestsResult, error) {
	command := NewCommand("ApplicationCache.getFramesWithManifests", nil)
	result := &applicationCache.GetFramesWithManifestsResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
GetManifestForFrame returns manifest URL for document in the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
*/
func (protocol *ApplicationCacheProtocol) GetManifestForFrame(
	params *applicationCache.GetManifestForFrameParams,
) (*applicationCache.GetManifestForFrameResult, error) {
	command := NewCommand("ApplicationCache.getManifestForFrame", params)
	result := &applicationCache.GetManifestForFrameResult{}
	protocol.Socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := json.Unmarshal(command.Result(), &result)
	return result, err
}

/*
OnApplicationCacheStatusUpdated adds a handler to the
ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
*/
func (protocol *ApplicationCacheProtocol) OnApplicationCacheStatusUpdated(
	callback func(event *applicationCache.StatusUpdatedEvent),
) {
	handler := NewEventHandler(
		"ApplicationCache.applicationCacheStatusUpdated",
		func(response *Response) {
			event := &applicationCache.StatusUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnNetworkStateUpdated adds a handler to the ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
*/
func (protocol *ApplicationCacheProtocol) OnNetworkStateUpdated(
	callback func(event *applicationCache.NetworkStateUpdatedEvent),
) {
	handler := NewEventHandler(
		"ApplicationCache.networkStateUpdated",
		func(response *Response) {
			event := &applicationCache.NetworkStateUpdatedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
