package network

import (
	"encoding/json"
	"fmt"
)

type cookieSameSiteEnum struct {
	Strict CookieSameSiteEnum
	Lax    CookieSameSiteEnum
}

var CookieSameSite = cookieSameSiteEnum{
	Strict: cookieSameSiteStrict,
	Lax:    cookieSameSiteLax,
}

/*
CookieSameSite represents the cookie's 'SameSite' status. Allowed values:
	- CookieSameSite.Strict
	- CookieSameSite.Lax

https://tools.ietf.org/html/draft-west-first-party-cookies

https://chromedevtools.github.io/devtools-protocol/tot/Network/#type-CookieSameSite
*/
type CookieSameSiteEnum int

/*
String implements Stringer
*/
func (enum CookieSameSiteEnum) String() string {
	return _cookieSameSiteEnums[enum]
}

/*
MarshalJSON implements json.Marshaler
*/
func (enum CookieSameSiteEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(enum.String())
}

/*
UnmarshalJSON implements json.Unmarshaler
*/
func (enum *CookieSameSiteEnum) UnmarshalJSON(bytes []byte) error {
	var err error
	var val string

	err = json.Unmarshal(bytes, &val)
	if nil != err {
		return err
	}

	for k, v := range _cookieSameSiteEnums {
		if v == val {
			*enum = k
			return nil
		}
	}

	return fmt.Errorf("%s is not a valid type value", bytes)
}

const (
	// cookieSameSiteStrict represents the "Strict" value.
	cookieSameSiteStrict CookieSameSiteEnum = iota + 1
	// cookieSameSiteLax represents the "Lax" value.
	cookieSameSiteLax
)

var _cookieSameSiteEnums = map[CookieSameSiteEnum]string{
	CookieSameSiteEnum(0): "",
	cookieSameSiteStrict:  "Strict",
	cookieSameSiteLax:     "Lax",
}
