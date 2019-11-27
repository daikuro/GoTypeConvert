package typeconv

import "fmt"

type StringValue struct {
	A     string // answer
	IsNil bool
}

var DefaultString = ""

func ToString(value interface{}, defaultValue ...string) *StringValue {
	if len(defaultValue) > 0 {
		return toStringD(value, defaultValue[0])
	}
	return toStringD(value, DefaultString)
}

func toStringD(value interface{}, defaultValue string) *StringValue {
	r := &StringValue{
		A:     defaultValue,
		IsNil: false,
	}
	if value == nil {
		r.IsNil = true
		return r
	}
	switch t := value.(type) {
	case string:
		r.A = t
		return r
	default:
		r.A = fmt.Sprint(value)
		return r
	}
	return r
}
