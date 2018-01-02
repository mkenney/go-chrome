package socket

import (
	"fmt"
	"sync"
)

/*
NewCommandMap creates and returns a pointer to a CommandMapper.
*/
func NewCommandMap() CommandMapper {
	return &commandMap{
		stack: make(map[int]Commander),
		mux:   &sync.Mutex{},
	}
}

/*
commandMap implements CommandMapper.
*/
type commandMap struct {
	stack map[int]Commander
	mux   *sync.Mutex
}

/*
Delete implements CommandMapper.
*/
func (stack *commandMap) Delete(id int) {
	delete(stack.stack, id)
}

/*
Get implements CommandMapper.
*/
func (stack *commandMap) Get(id int) (Commander, error) {
	command, ok := stack.stack[id]
	if !ok {
		return nil, fmt.Errorf("Command %d not found", id)
	}
	return command, nil
}

/*
Lock implements CommandMapper.
*/
func (stack *commandMap) Lock() {
	stack.mux.Lock()
}

/*
Set implements CommandMapper.
*/
func (stack *commandMap) Set(cmd Commander) {
	stack.stack[cmd.ID()] = cmd
}

/*
Unlock implements CommandMapper.
*/
func (stack *commandMap) Unlock() {
	stack.mux.Unlock()
}
