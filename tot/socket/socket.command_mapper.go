package socket

import (
	"fmt"
	"sync"
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
	stack map[int]Commander
	mux   *sync.Mutex
}

/*
Delete removes a command from the stack.

Delete is a CommandMapper implementation.
*/
func (stack *CommandMap) Delete(id int) {
	delete(stack.stack, id)
}

/*
Get retrieves a command from the stack.

Get is a CommandMapper implementation.
*/
func (stack *CommandMap) Get(id int) (Commander, error) {
	command, ok := stack.stack[id]
	if !ok {
		return nil, fmt.Errorf("Command %d not found", id)
	}
	return command, nil
}

/*
Lock locks the sync mutex.

Lock is a CommandMapper implementation.
*/
func (stack *CommandMap) Lock() {
	stack.mux.Lock()
}

/*
Set sets a command in the stack.

Set is a CommandMapper implementation.
*/
func (stack *CommandMap) Set(cmd Commander) {
	stack.stack[cmd.ID()] = cmd
}

/*
Unlock unlocks the sync mutex.

Unlock is a CommandMapper implementation.
*/
func (stack *CommandMap) Unlock() {
	stack.mux.Unlock()
}
