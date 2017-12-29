package chrome

import (
	"fmt"
	"strings"
)

/*
Args contains CLI arguments to the Chromium executable.
*/
type Args map[string][]interface{}

/*
Get implements Commander
*/
func (args Args) Get(arg string) (values []interface{}, err error) {
	if !args.Has(arg) {
		err = fmt.Errorf("The specified argument '%s' does not exist", arg)
	} else {
		values = args[arg]
	}
	return values, err
}

/*
Has implements Commander
*/
func (args Args) Has(arg string) bool {
	_, ok := args[arg]
	return ok
}

/*
List implements Commander
*/
func (args Args) List() []string {
	var list []string

	for arg, vals := range args {
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
Set implements Commander
*/
func (args Args) Set(arg string, values []interface{}) (err error) {
	if nil == values {
		values = make([]interface{}, 0)
	}
	args[arg] = values

	for k, value := range values {
		if nil != value {
			switch value.(type) {
			case int:
				args[arg][k] = value.(int)
			case string:
				args[arg][k] = value.(string)
			default:
				panic(fmt.Sprintf("Invalid data type %v for argument %s", value, arg))
			}
		}
	}

	return nil
}

/*
String implements Commander
*/
func (args Args) String() string {
	return strings.Join(args.List(), " ")
}
