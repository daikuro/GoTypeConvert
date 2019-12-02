package typeconv

import (
	"strconv"
)

type IntValue struct {
	A     int // answer
	Error error
}

var DefaultInt = 0

func ToInt(value interface{}, defaultValue ...int) *IntValue {
	r := &IntValue{
		A: DefaultInt,
	}
	if len(defaultValue) > 0 {
		r.A = defaultValue[0]
	}
	switch t := value.(type) {
	case int:
		r.A = t
		return r
	case int8:
		r.A = int(t)
	case int16:
		r.A = int(t)
	case int32:
		r.A = int(t)
	case int64:
		r.A = int(t)
	case string:
		r.A, r.Error = strconv.Atoi(t)
	case float32:
		r.A = int(t)
	case float64:
		r.A = int(t)
	case bool:
		if t {
			r.A = 1
		} else {
			r.A = 0
		}
	}
	return r
}
