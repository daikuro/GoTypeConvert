package typeconv

import "strings"

type BoolValue struct {
	A     bool // answer
	Error error
}

var DefaultBool = false

func ToBool(value interface{}) *BoolValue {
	return ToBoold(value, DefaultBool)
}

func ToBoold(value interface{}, defaultValue bool) *BoolValue {
	r := &BoolValue{
		A: defaultValue,
	}
	switch t := value.(type) {
	case bool:
		r.A = t
	case string:
		switch strings.ToLower(t) {
		case "true":
			r.A = true
		case "false":
			r.A = false
		}
	case int:
		switch t {
		case 1:
			r.A = true
		case 0:
			r.A = false
		}
	}
	return r
}
