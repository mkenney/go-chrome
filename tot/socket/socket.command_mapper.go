package socket

import (
	"fmt"
	"sync"
)

/*
NewCommandMap creates and returns a pointer to a CommandMapper.
*/
func NewCommandMap() *CommandMap {
	return &CommandMap{
		stack: make(map[int]Commander),
		mux:   &sync.Mutex{},
	}
}

/*
CommandMap implements CommandMapper.
*/
type CommandMap struct {
	stack map[int]Commander
	mux   *sync.Mutex
}

/*
Delete implements CommandMapper.
*/
func (stack *CommandMap) Delete(id int) {
	delete(stack.stack, id)
}

/*
Get implements CommandMapper.
*/
func (stack *CommandMap) Get(id int) (Commander, error) {
	command, ok := stack.stack[id]
	if !ok {
		return nil, fmt.Errorf("Command %d not found", id)
	}
	return command, nil
}

/*
Lock implements CommandMapper.
*/
func (stack *CommandMap) Lock() {
	stack.mux.Lock()
}

/*
Set implements CommandMapper.
*/
func (stack *CommandMap) Set(cmd Commander) {
	stack.stack[cmd.ID()] = cmd
}

/*
Unlock implements CommandMapper.
*/
func (stack *CommandMap) Unlock() {
	stack.mux.Unlock()
}
