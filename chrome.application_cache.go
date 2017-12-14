package chrome

import (
	application_cache "app/chrome/application_cache"
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
ApplicationCache - https://chromedevtools.github.io/devtools-protocol/tot/ApplicationCache/
EXPERIMENTAL
*/
type ApplicationCache struct{}

/*
Enable enables application cache domain notifications.
*/
func (ApplicationCache) Enable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "ApplicationCache.enable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
Disable disables application cache domain notifications.
*/
func (ApplicationCache) Disable(
	socket *Socket,
) error {
	command := &protocol.Command{
		Method: "ApplicationCache.disable",
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetApplicationCacheForFrame returns relevant application cache data for the document
in given frame.
*/
func (ApplicationCache) GetApplicationCacheForFrame(
	socket *Socket,
	params *application_cache.GetApplicationCacheForFrameParams,
) (application_cache.GetApplicationCacheForFrameResult, error) {
	command := &protocol.Command{
		Method: "ApplicationCache.getManifestForFrame",
		Params: params,
	}
	result := application_cache.GetApplicationCacheForFrameResult{}
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
*/
func (ApplicationCache) GetFramesWithManifests(
	socket *Socket,
) (application_cache.GetFramesWithManifestsResult, error) {
	command := &protocol.Command{
		Method: "ApplicationCache.getFramesWithManifests",
	}
	result := application_cache.GetFramesWithManifestsResult{}
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
*/
func (ApplicationCache) GetManifestForFrame(
	socket *Socket,
	params *application_cache.GetManifestForFrameParams,
) (application_cache.GetManifestForFrameResult, error) {
	command := &protocol.Command{
		Method: "ApplicationCache.getManifestForFrame",
		Params: params,
	}
	result := application_cache.GetManifestForFrameResult{}
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
OnApplicationCacheStatusUpdated adds a handler to the ApplicationCache.applicationCacheStatusUpdated
event.
*/
func (ApplicationCache) OnApplicationCacheStatusUpdated(
	socket *Socket,
	callback func(event *application_cache.ApplicationCacheStatusUpdatedEvent),
) {
	handler := protocol.NewEventHandler(
		"ApplicationCache.applicationCacheStatusUpdated",
		func(name string, params []byte) {
			event := &application_cache.ApplicationCacheStatusUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
