package typeconv

import (
	"strconv"
)

type IntValue struct {
	A     int // answer
	Error error
}

var DefaultInt = 0

func ToInt(value interface{}) *IntValue {
	return ToIntd(value, DefaultInt)
}

func ToIntd(value interface{}, defaultValue int) *IntValue {
	r := &IntValue{
		A: defaultValue,
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
	}
	return r
}
