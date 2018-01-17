package debugger

import (
	"encoding/json"
	"fmt"
)

type scopeTypeEnum struct {
	Global  ScopeTypeEnum
	Local   ScopeTypeEnum
	With    ScopeTypeEnum
	Closure ScopeTypeEnum
	Catch   ScopeTypeEnum
	Block   ScopeTypeEnum
	Script  ScopeTypeEnum
	Eval    ScopeTypeEnum
	Module  ScopeTypeEnum
}

var ScopeType = scopeTypeEnum{
	Global:  ScopeTypeGlobal,
	Local:   ScopeTypeLocal,
	With:    ScopeTypeWith,
	Closure: ScopeTypeClosure,
	Catch:   ScopeTypeCatch,
	Block:   ScopeTypeBlock,
	Script:  ScopeTypeScript,
	Eval:    ScopeTypeEval,
	Module:  ScopeTypeModule,
}

/*
Scope type. Allowed values:
	- global
	- local
	- with
	- closure
	- catch
	- block
	- script
	- eval
	- module

https://chromedevtools.github.io/devtools-protocol/tot/Debugger/#type-Scope
*/
type ScopeTypeEnum int

/*
String implements Stringer
*/
func (enum ScopeTypeEnum) String() string {
	return _scopeTypeEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum ScopeTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *ScopeTypeEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _scopeTypeEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid scopeType value", bytes)
}

const (
	// ScopeTypeGlobal represents the "global" value.
	ScopeTypeGlobal ScopeTypeEnum = iota + 1
	// ScopeTypeLocal represents the "local" value.
	ScopeTypeLocal
	// ScopeTypeWith represents the "with" value.
	ScopeTypeWith
	// ScopeTypeClosure represents the "closure" value.
	ScopeTypeClosure
	// ScopeTypeCatch represents the "catch" value.
	ScopeTypeCatch
	// ScopeTypeBlock represents the "block" value.
	ScopeTypeBlock
	// ScopeTypeScript represents the "script" value.
	ScopeTypeScript
	// ScopeTypeEval represents the "eval" value.
	ScopeTypeEval
	// ScopeTypeModule represents the "module" value.
	ScopeTypeModule
)

var _scopeTypeEnums = map[ScopeTypeEnum]string{
	ScopeTypeGlobal:  "global",
	ScopeTypeLocal:   "local",
	ScopeTypeWith:    "with",
	ScopeTypeClosure: "closure",
	ScopeTypeCatch:   "catch",
	ScopeTypeBlock:   "block",
	ScopeTypeScript:  "script",
	ScopeTypeEval:    "eval",
	ScopeTypeModule:  "module",
}
