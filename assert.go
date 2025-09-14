package assert

import (
	"fmt"
	"log"
)

var Debug = true

func Catch(fn func(any)) {
	e := recover()
	log.Printf("assert.Catch(%T)", e)
	if fn != nil {
		fn(e)
	}
}

func Z(w any, s any, d ...any) {
	True(w == nil, s, d...)
	switch w.(type) {
	case string:
		True(w == "", s, d...)
	case bool:
		True(w == false, s, d...)
	case int:
		True(w == 0, s, d...)
	default:
		True(false, "unknown type of %T", w)
	}
}

func Ok(w any, s any, d ...any) {
	True(w != nil, s, d...)
	switch w.(type) {
	case string:
		True(w != "", s, d...)
	case bool:
		True(w == true, s, d...)
	case int:
		True(w != 0, s, d...)
	default:
	}
}

type Assertion struct {
	Message string
}

func (a Assertion) Error() string {
	return a.Message
}

func NewAssertion(msg string) Assertion {
	return Assertion{}
}

func T(c bool, err error) {
	if Debug {
		log.Printf("assert:T(%t, %T)", c, err)
	}
	if c {
		return
	}

	panic(err)
}

func True(c bool, s any, d ...any) {
	var m string

	switch s.(type) {
	case string:
		m = fmt.Sprintf(s.(string), d...)
	default:
		m = fmt.Sprintf("%T", s)
	}

	if c {
		if Debug {
			log.Printf("pass: %s\n", m)
		}
		return
	}

	log.Printf("%T", m)
	panic(m)
}
