package socket

import (
	"encoding/json"
	"fmt"
	"testing"

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
	cmd := NewCommand("Some.method", params)

	if "Some.method" != cmd.Method() {
		t.Errorf("Expected 'Some.method', got '%s'", cmd.Method())
	}

	if 1 != cmd.Params().(mockParams).Value {
		t.Errorf("Expected 1, got '%d'", cmd.Params().(mockParams).Value)
	}
}

func TestCommanderNextID(t *testing.T) {
	// Ids generated in order
	testMux.Lock()
	start := _commandID
	for a := start + 1; a <= start+1000; a++ {
		cmd := NewCommand("Some.method", nil)
		if a != cmd.ID() {
			t.Errorf("nextID() failed to generate consecutive IDs - expeted %d, received %d", a, cmd.ID())
		}
	}
	testMux.Unlock()

	// Ids generated safely
	id := make(chan int)
	for a := 0; a <= 1000; a++ {
		go func(id chan int) {
			cmd := NewCommand("Some.method", nil)
			log.Debugf("Sending id %d", cmd.ID())
			id <- cmd.ID()
		}(id)
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

func TestCommanderDone(t *testing.T) {
	cmd := NewCommand("Some.method", nil)
	for a := 0; a < 1000; a++ {
		result := mockResult{
			Message: "Mock Message",
			Data: map[string]int{
				"value": a,
			},
		}
		cmd.WaitGroup().Add(1)
		go func() {
			resultData, err := json.Marshal(result)
			cmd.Done(resultData, err)
			if string(resultData) != string(cmd.Result()) {
				t.Errorf("Expected '%s', got '%s'", resultData, cmd.Result())
			}
		}()
		cmd.WaitGroup().Wait()
	}
}

func TestCommanderError(t *testing.T) {
	cmd := NewCommand("Some.method", nil)
	result := mockResult{
		Message: "Mock Message",
		Data: map[string]int{
			"value": 1,
		},
	}
	resultBytes, err := json.Marshal(result)
	cmd.WaitGroup().Add(1)
	cmd.Done(resultBytes, err)
	if nil != cmd.Error() {
		t.Errorf("Expected nil, got error: '%s'", cmd.Error().Error())
	}

	cmd = NewCommand("Some.method", nil)
	cmd.WaitGroup().Add(1)
	cmd.Done(nil, fmt.Errorf("This is an error"))
	if nil == cmd.Error() {
		t.Errorf("Expected error, got nil")
	}
	if "This is an error" != cmd.Error().Error() {
		t.Errorf("Expected error 'This is an error', got error: '%s'", cmd.Error().Error())
	}
}
