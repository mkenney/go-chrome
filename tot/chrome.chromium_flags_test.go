package chrome

import (
	"testing"
)

func TestChromiumFlagsGet(t *testing.T) {
	flags := &Flags{}

	value, err := flags.Get("test-arg")
	if nil == err {
		t.Errorf("Expected error, received nil")
	}
	if nil != value {
		t.Errorf("Expected nil, received %v", value)
	}
}

func TestChromiumFlagsHas(t *testing.T) {
	flags := &Flags{}

	has := flags.Has("test-arg")
	if has {
		t.Errorf("Expected false, received true")
	}
}

func TestChromiumFlagsList(t *testing.T) {
	flags := &Flags{}
	list := flags.List()
	if nil != list {
		t.Errorf("expected nil, received %T: %+v", list, list)
	}

	flags = &Flags{
		"test-1": nil,
		"test-2": "string",
		"test-3": 1,
	}

	list = flags.List()
	if nil == list {
		t.Errorf("Expected argument list, received nil")
	}
	if 3 != len(list) {
		t.Errorf("Expected 3 arguments, received %d", len(list))
	}
	if "--test-1" != list[0] {
		t.Errorf("Expected '--test-1', received %s", list[0])
	}
	if "--test-2=string" != list[1] {
		t.Errorf("Expected '--test-2=string', received '%s'", list[1])
	}
	if "--test-3=1" != list[2] {
		t.Errorf("Expected '--test-3=1', received '%s'", list[2])
	}
}

func TestChromiumFlagsSet(t *testing.T) {
	var err error
	flags := &Flags{}

	err = flags.Set("test-1", nil)
	if nil != err {
		t.Errorf("Expected nil, received error '%s'", err.Error())
	}

	err = flags.Set("test-2", "string")
	if nil != err {
		t.Errorf("Expected nil, received error '%s'", err.Error())
	}

	err = flags.Set("test-3", 1)
	if nil != err {
		t.Errorf("Expected nil, received error '%s'", err.Error())
	}

	err = flags.Set("test-4", false)
	if nil == err {
		t.Errorf("Expected error, received nil")
	}

	args := flags.String()
	if "" == args {
		t.Errorf("Expected argument list, received empty string")
	}
	if "--test-1 --test-2=string --test-3=1" != args {
		t.Errorf("Expected '--test-1 --test-2=string --test-3=1', received '%s'", args)
	}
}

func TestChromiumFlagsString(t *testing.T) {
	flags := &Flags{
		"test-1": nil,
		"test-2": "string",
		"test-3": 1,
	}

	args := flags.String()
	if "" == args {
		t.Errorf("Expected argument list, received empty string")
	}
	if "--test-1 --test-2=string --test-3=1" != args {
		t.Errorf("Expected '--test-1 --test-2=string --test-3=1', received '%s'", args)
	}
}
