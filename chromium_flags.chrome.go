package chrome

import (
	"fmt"
	"strings"
)

/*
Flags contains CLI arguments to the Chromium executable.
*/
type Flags map[string][]interface{}

/*
Get implements ChromiumFlags
*/
func (flags Flags) Get(arg string) (values []interface{}, err error) {
	if !flags.Has(arg) {
		err = fmt.Errorf("The specified argument '%s' does not exist", arg)
	} else {
		values = flags[arg]
	}
	return values, err
}

/*
Has implements ChromiumFlags
*/
func (flags Flags) Has(arg string) bool {
	_, ok := flags[arg]
	return ok
}

/*
List implements ChromiumFlags
*/
func (flags Flags) List() []string {
	var list []string

	for arg, vals := range flags {
		if nil == vals {
			list = append(list, fmt.Sprintf("--%s", arg))
		} else {
			for _, val := range vals {
				switch val.(type) {
				case int:
					arg = fmt.Sprintf("--%s=%d", arg, val)
				case string:
					arg = fmt.Sprintf("--%s=%s", arg, val)
				default:
					panic(fmt.Sprintf("Invalid value data type %v", val))
				}
				list = append(list, arg)
			}
		}
	}

	return list
}

/*
Set implements ChromiumFlags
*/
func (flags Flags) Set(arg string, values []interface{}) (err error) {
	if nil == values {
		values = make([]interface{}, 0)
	}
	flags[arg] = values

	for k, value := range values {
		if nil != value {
			switch value.(type) {
			case int:
				flags[arg][k] = value.(int)
			case string:
				flags[arg][k] = value.(string)
			default:
				panic(fmt.Sprintf("Invalid data type %v for argument %s", value, arg))
			}
		}
	}

	return nil
}

/*
String implements ChromiumFlags
*/
func (flags Flags) String() string {
	return strings.Join(flags.List(), " ")
}
