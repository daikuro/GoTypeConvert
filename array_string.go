package typeconv

import "fmt"

type ArrayStringValue struct {
	A []string // value
	T string   // struct name
}

var DefaultArrayString = &ArrayStringValue{
	A: []string{},
}

func ToArrayString(value interface{}, defaultValue ...[]string) *ArrayStringValue {
	if defaultValue != nil && len(defaultValue) > 0 {
		d := &ArrayStringValue{
			A: defaultValue[0],
		}
		return toArrayString(value, d)
	}
	return toArrayString(value, DefaultArrayString)
}

func toArrayString(value interface{}, defaultValue *ArrayStringValue) *ArrayStringValue {
	r := &ArrayStringValue{
		A: []string{},
	}
	switch t := value.(type) {
	case []string:
		r.A = t
	case []interface{}:
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = toStringD(value, "").A
		}
	default:
		r = defaultValue
	}
	r.T = fmt.Sprintf("%T", value)
	return r
}
