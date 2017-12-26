package chrome

/*
LoadEvent is the load event data
*/
type LoadEvent struct {
	Timestamp float64 `json:"timestamp"`
}

///*
//OnLoadEvent fires when the page loads
//*/
//func OnLoadEvent(socket *Socket, fn func(evt *LoadEvent)) {
//	socket.AddEventHandler("Page.loadEventFired", func(name string, params []byte) {
//		evt := &LoadEvent{}
//		if err := json.Unmarshal(params, evt); err != nil {
//			log.Error(err)
//		} else {
//			fn(evt)
//		}
//	})
//}
