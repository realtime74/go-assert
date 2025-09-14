package assert

import (
	"fmt"
	"log"
)

var debug = true

type Assert struct{}

func asserte(c any, s string, d ...any) {
	True(c != nil, "is nil")
	switch c.(type) {
	case string:
		True(c != "", s, d...)
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

func I(p any, s string, d ...any) {
	True(p != nil, s, d...)
}
func P(p *any, s string, d ...any) {
	True(p != nil, s, d...)
}

func True(c bool, s any, d ...any) {
	var m string

	switch s.(type) {
	case string:
		m = fmt.Sprintf(s.(string), d...)
	default:
	}

	if c {
		if debug {
			log.Printf("pass: %s\n", m)
		}
		return
	}

	log.Printf("%T", m)
	panic(m)
}

func blablue() string {
	return ""
}

func do_bad() {
	True(false, "really bad\n")
}

func Catch(f func(err any)) {
	a := recover()
	fmt.Printf("got you: %s", a)
	f(a)
}

func catch_bad() {
	func() {
		defer Catch(func(_ any) {})
		do_bad()
	}()

	func() {
		defer func() {
			fmt.Printf("gotyoux: %s", recover())
		}()

		do_bad()

	}()

	fmt.Printf("still alive\n")
}

func _catch(s string) {
	a := recover()
	fmt.Printf("catched[%s]: %s\n", a, s)
}

func try_something() {
	defer _catch("from: trysomething")

	do_bad()

}
