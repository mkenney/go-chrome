package socket

import (
	"encoding/json"

	"github.com/mkenney/go-chrome/tot/application/cache"
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
func (protocol *ApplicationCacheProtocol) Enable() <-chan *cache.EnableResult {
	resultChan := make(chan *cache.EnableResult)
	command := NewCommand(protocol.Socket, "ApplicationCache.enable", nil)
	result := &cache.EnableResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetForFrame returns relevant application cache data for the document
in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
func (protocol *ApplicationCacheProtocol) GetForFrame(
	params *cache.GetForFrameParams,
) <-chan *cache.GetForFrameResult {
	resultChan := make(chan *cache.GetForFrameResult)
	command := NewCommand(protocol.Socket, "ApplicationCache.getApplicationCacheForFrame", params)
	result := &cache.GetForFrameResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetFramesWithManifests returns array of frame identifiers with manifest urls for
each frame containing a document associated with some application cache.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
*/
func (protocol *ApplicationCacheProtocol) GetFramesWithManifests() <-chan *cache.GetFramesWithManifestsResult {
	resultChan := make(chan *cache.GetFramesWithManifestsResult)
	command := NewCommand(protocol.Socket, "ApplicationCache.getFramesWithManifests", nil)
	result := &cache.GetFramesWithManifestsResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
GetManifestForFrame returns manifest URL for document in the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
*/
func (protocol *ApplicationCacheProtocol) GetManifestForFrame(
	params *cache.GetManifestForFrameParams,
) <-chan *cache.GetManifestForFrameResult {
	resultChan := make(chan *cache.GetManifestForFrameResult)
	command := NewCommand(protocol.Socket, "ApplicationCache.getManifestForFrame", params)
	result := &cache.GetManifestForFrameResult{}

	go func() {
		response := <-protocol.Socket.SendCommand(command)
		if nil != response.Error && 0 != response.Error.Code {
			result.Err = response.Error
		} else {
			result.Err = json.Unmarshal(response.Result, &result)
		}
		resultChan <- result
		close(resultChan)
	}()

	return resultChan
}

/*
OnApplicationCacheStatusUpdated adds a handler to the
ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
*/
func (protocol *ApplicationCacheProtocol) OnApplicationCacheStatusUpdated(
	callback func(event *cache.StatusUpdatedEvent),
) {
	handler := NewEventHandler(
		"ApplicationCache.applicationCacheStatusUpdated",
		func(response *Response) {
			event := &cache.StatusUpdatedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}

/*
OnNetworkStateUpdated adds a handler to the ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-networkStateUpdated
*/
func (protocol *ApplicationCacheProtocol) OnNetworkStateUpdated(
	callback func(event *cache.NetworkStateUpdatedEvent),
) {
	handler := NewEventHandler(
		"ApplicationCache.networkStateUpdated",
		func(response *Response) {
			event := &cache.NetworkStateUpdatedEvent{}
			json.Unmarshal([]byte(response.Params), event)
			if nil != response.Error && 0 != response.Error.Code {
				event.Err = response.Error
			}
			callback(event)
		},
	)
	protocol.Socket.AddEventHandler(handler)
}
