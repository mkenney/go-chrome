package chrome

import (
	"fmt"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
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

	orderedFlags := []string{}
	for arg := range flags {
		orderedFlags = append(orderedFlags, arg)
	}
	sort.Strings(orderedFlags)

	for _, arg := range orderedFlags {
		vals, err := flags.Get(arg)
		if nil != err {
			log.Fatal(err)
		}
		if nil == vals || 0 == len(vals) {
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
		if _, ok := flags[arg]; !ok {
			flags[arg] = nil
		}
	}

	for _, value := range values {
		if nil != value {
			switch value.(type) {
			case int:
				flags[arg] = append(flags[arg], value.(int))
			case string:
				flags[arg] = append(flags[arg], value.(string))
			default:
				return fmt.Errorf("Invalid data type %v for argument %s", value, arg)
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
