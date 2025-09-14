package assert

import (
	"fmt"
	"log"
)

var Debug = true

func Catch(fn func(err any)) {
	err := recover()
	if fn != nil {
		fn(err)
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

func True(c bool, s any, d ...any) {
	var m string

	switch s.(type) {
	case string:
		m = fmt.Sprintf(s.(string), d...)
	default:
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
