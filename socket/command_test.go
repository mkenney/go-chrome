package socket

import (
	"testing"
)

func TestMutex(t *testing.T) {
	//	tmp, _ := newMock("https://www.example.com/")
	//	mockSocket := interface{}(tmp).(*socket)
	//
	//	set := func(a int) {
	//		mockSocket.Lock()
	//		mockSocket.commandID = a
	//		mockSocket.Unlock()
	//	}
	//
	//	get := func(ch chan int) {
	//		mockSocket.Lock()
	//		mockSocket.Unlock()
	//		ch <- mockSocket.commandID
	//	}
	//	for a := 0; a <= 10; a++ {
	//		for a := 0; a <= 1000; a++ {
	//			go set(a)
	//		}
	//
	//		time.Sleep(time.Second * 3)
	//		ch := make(chan int)
	//		go get(ch)
	//		if 1000 != <-ch {
	//			t.Errorf("Socketer mutexes to not appear to be working correctly: 1000 expected, %d received", mockSocket.commandID)
	//			return
	//		}
	//	}
}
