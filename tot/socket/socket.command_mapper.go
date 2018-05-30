package socket

import (
	"fmt"
	"sync"

	errs "github.com/mkenney/go-errors"
)

/*
NewCommandMap creates and returns a pointer to a struct that implements the
CommandMapper interface.
*/
func NewCommandMap() *CommandMap {
	return &CommandMap{
		stack: make(map[int]Commander),
		mux:   &sync.Mutex{},
	}
}

/*
CommandMap provides a CommandMapper interface implementation for managing the
command stack.
*/
type CommandMap struct {
	mux   *sync.Mutex
	stack map[int]Commander
}

/*
Delete removes a command from the stack.

Delete is a CommandMapper implementation.
*/
func (stack *CommandMap) Delete(id int) {
	stack.mux.Lock()
	delete(stack.stack, id)
	stack.mux.Unlock()
}

/*
Get retrieves a command from the stack.

Get is a CommandMapper implementation.
*/
func (stack *CommandMap) Get(id int) (Commander, error) {
	stack.mux.Lock()
	command, ok := stack.stack[id]
	stack.mux.Unlock()
	if !ok {
		return nil, errs.New(fmt.Sprintf("Command %d not found", id))
	}
	return command, nil
}

/*
Set sets a command in the stack.

Set is a CommandMapper implementation.
*/
func (stack *CommandMap) Set(cmd Commander) {
	stack.mux.Lock()
	stack.stack[cmd.ID()] = cmd
	stack.mux.Unlock()
}
