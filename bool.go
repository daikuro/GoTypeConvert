package typeconv

import "strings"

type BoolValue struct {
	A       bool // answer
	NoValue bool
}

var DefaultBool = false

func ToBool(value interface{}) *BoolValue {
	return ToBoold(value, DefaultBool)
}

func ToBoold(value interface{}, defaultValue bool) *BoolValue {
	r := &BoolValue{
		A:       defaultValue,
		NoValue: true,
	}
	switch t := value.(type) {
	case bool:
		r.A = t
		r.NoValue = false
	case string:
		switch strings.ToLower(t) {
		case "true":
			r.A = true
			r.NoValue = false
		case "false":
			r.A = false
			r.NoValue = false
		}
	case int:
		switch t {
		case 1:
			r.A = true
			r.NoValue = false
		case 0:
			r.A = false
			r.NoValue = false
		}
	}
	return r
}
