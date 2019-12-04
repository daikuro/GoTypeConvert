package typeconv

import (
	"fmt"
	"io"
	"io/ioutil"
)

type StringValue struct {
	A     string // answer
	IsNil bool
	Error error
}

var DefaultString = ""

func ToString(value interface{}, defaultValue ...string) (r *StringValue) {
	r = &StringValue{
		A:     DefaultString,
		IsNil: false,
	}
	if len(defaultValue) > 0 {
		r.A = defaultValue[0]
	}

	if value == nil {
		r.IsNil = true
		return r
	}
	switch t := value.(type) {
	case string:
		r.A = t
		return r
	case []byte:
		r.A = string(t)
		return r
	case io.Reader:
		b, err := ioutil.ReadAll(t)
		if err != nil {
			r.Error = err
		} else {
			r.A = string(b)
		}
		return
	default:
		r.A = fmt.Sprint(value)
		return r
	}
}
