package protocol

import (
	"encoding/json"

	applicationCache "github.com/mkenney/go-chrome/protocol/application_cache"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
ApplicationCache - https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/
EXPERIMENTAL
*/
type ApplicationCache struct{}

/*
Enable enables application cache domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-enable
*/
func (ApplicationCache) Enable(
	socket *sock.Socket,
) error {
	command := &sock.Command{
		Method: "ApplicationCache.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetForFrame returns relevant application cache data for the document
in given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getApplicationCacheForFrame
*/
func (ApplicationCache) GetForFrame(
	socket *sock.Socket,
	params *applicationCache.GetForFrameParams,
) (applicationCache.GetForFrameResult, error) {
	command := &sock.Command{
		Method: "ApplicationCache.getApplicationCacheForFrame",
		Params: params,
	}
	result := applicationCache.GetForFrameResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetFramesWithManifests returns array of frame identifiers with manifest urls for each frame
containing a document associated with some application cache.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getFramesWithManifests
*/
func (ApplicationCache) GetFramesWithManifests(
	socket *sock.Socket,
) (applicationCache.GetFramesWithManifestsResult, error) {
	command := &sock.Command{
		Method: "ApplicationCache.getFramesWithManifests",
	}
	result := applicationCache.GetFramesWithManifestsResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
GetManifestForFrame returns manifest URL for document in the given frame.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#method-getManifestForFrame
*/
func (ApplicationCache) GetManifestForFrame(
	socket *sock.Socket,
	params *applicationCache.GetManifestForFrameParams,
) (applicationCache.GetManifestForFrameResult, error) {
	command := &sock.Command{
		Method: "ApplicationCache.getManifestForFrame",
		Params: params,
	}
	result := applicationCache.GetManifestForFrameResult{}
	socket.SendCommand(command)

	if nil != command.Err {
		return result, command.Err
	}

	if nil != command.Result {
		resultData, err := json.Marshal(command.Result)
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Err
}

/*
OnApplicationCacheStatusUpdated adds a handler to the ApplicationCache.StatusUpdated event.

https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/#event-applicationCacheStatusUpdated
*/
func (ApplicationCache) OnApplicationCacheStatusUpdated(
	socket *sock.Socket,
	callback func(event *applicationCache.StatusUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ApplicationCache.applicationCacheStatusUpdated",
		func(name string, params []byte) {
			event := &applicationCache.StatusUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
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
func (ApplicationCache) OnNetworkStateUpdated(
	socket *sock.Socket,
	callback func(event *applicationCache.StatusUpdatedEvent),
) {
	handler := sock.NewEventHandler(
		"ApplicationCache.networkStateUpdated",
		func(name string, params []byte) {
			event := &applicationCache.StatusUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
