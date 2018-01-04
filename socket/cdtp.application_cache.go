package socket

import (
	"encoding/json"

	applicationCache "github.com/mkenney/go-chrome/cdtp/application_cache"
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
func (protocol *ApplicationCacheProtocol) Enable() chan *applicationCache.EnableResult {
	resultChan := make(chan *applicationCache.EnableResult)
	command := NewCommand(protocol.Socket, "ApplicationCache.enable", nil)
	result := &applicationCache.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetForFrame returns relevant application cache data for the document
in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
func (protocol *ApplicationCacheProtocol) GetForFrame(
	params *applicationCache.GetForFrameParams,
) chan *applicationCache.GetForFrameResult {
	resultChan := make(chan *applicationCache.GetForFrameResult)
	command := NewCommand(protocol.Socket, "ApplicationCache.getApplicationCacheForFrame", params)
	result := &applicationCache.GetForFrameResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetFramesWithManifests returns array of frame identifiers with manifest urls for
each frame containing a document associated with some application cache.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
*/
func (protocol *ApplicationCacheProtocol) GetFramesWithManifests() chan *applicationCache.GetFramesWithManifestsResult {
	resultChan := make(chan *applicationCache.GetFramesWithManifestsResult)
	command := NewCommand(protocol.Socket, "ApplicationCache.getFramesWithManifests", nil)
	result := &applicationCache.GetFramesWithManifestsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
}

/*
GetManifestForFrame returns manifest URL for document in the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
*/
func (protocol *ApplicationCacheProtocol) GetManifestForFrame(
	params *applicationCache.GetManifestForFrameParams,
) chan *applicationCache.GetManifestForFrameResult {
	resultChan := make(chan *applicationCache.GetManifestForFrameResult)
	command := NewCommand(protocol.Socket, "ApplicationCache.getManifestForFrame", params)
	result := &applicationCache.GetManifestForFrameResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.CDTPError = response.Error
		} else {
			result.CDTPError = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
	}()

	return resultChan
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
