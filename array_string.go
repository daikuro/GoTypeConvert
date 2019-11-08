package typeconv

import "fmt"

type ArrayStringValue struct {
	A []string // value
	T string   // struct name
}

var DefaultArrayString = &ArrayStringValue{
	A: []string{},
}

func ToArrayString(value interface{}) *ArrayStringValue {
	return ToArrayStringd(value, DefaultArrayString)
}

func ToArrayStringd(value interface{}, defaultValue *ArrayStringValue) *ArrayStringValue {
	r := &ArrayStringValue{
		A: []string{},
	}
	switch t := value.(type) {
	case []string:
		r.A = t
	case []interface{}:
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToStringd(value, "")
		}
	default:
		r = defaultValue
	}
	r.T = fmt.Sprintf("%T", value)
	return r
}
