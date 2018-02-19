package socket

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/socket"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, err := log.ParseLevel("debug")
	if nil == err {
		log.SetLevel(level)
	}
}

type mockParams struct {
	Value int `json:"value"`
}

type mockResult struct {
	Message string         `json:"message"`
	Data    map[string]int `json:"data"`
}

func TestNewCommander(t *testing.T) {
	params := mockParams{
		Value: 1,
	}
	socketURL, _ := url.Parse("https://www.example.com/mock")
	cmd := socket.NewCommand(socket.NewMock(socketURL), "Some.method", params)

	if "Some.method" != cmd.Method() {
		t.Errorf("Expected 'Some.method', got '%s'", cmd.Method())
	}

	if 1 != cmd.Params().(mockParams).Value {
		t.Errorf("Expected 1, got '%d'", cmd.Params().(mockParams).Value)
	}
}

func TestCommanderNextID(t *testing.T) {
	// Ids generated safely
	id := make(chan int)
	socketURL, _ := url.Parse("https://www.example.com/mock")
	mockSocket := socket.NewMock(socketURL)
	for a := 0; a <= 1000; a++ {
		go func() {
			cmd := socket.NewCommand(mockSocket, "Some.method", nil)
			log.Debugf("Sending id %d", cmd.ID())
			id <- cmd.ID()
		}()
	}
	results := make(map[int]bool)
	recCount := 0
	for nextID := range id {
		recCount++
		log.Debugf("Received ID %d", nextID)
		if _, ok := results[nextID]; !ok {
			results[nextID] = true
		} else {
			t.Errorf("nextID() failed to generate safe IDs - duplicate ID received %d", nextID)
		}
		if recCount >= 1000 {
			break
		}
	}
}

func TestCommanderError(t *testing.T) {
	params := mockParams{
		Value: 1,
	}
	socketURL, _ := url.Parse("https://www.example.com/mock")
	cmd := socket.NewCommand(socket.NewMock(socketURL), "Some.method", params)
	err := fmt.Errorf("test error")

	cmd.SetError(err)
	if nil == cmd.Error() {
		t.Errorf("Expected nil, got error: '%s'", cmd.Error().Error())
	}
	if err.Error() != cmd.Error().Error() {
		t.Errorf("Expected '%s', got '%s'", err.Error(), cmd.Error().Error())
	}
}
