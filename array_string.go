package typeconv

import "fmt"

type ArrayStringValue struct {
	A []string // value
	T string   // struct name
}

func ToArrayString(value interface{}, defaultValue ...[]string) *ArrayStringValue {
	r := &ArrayStringValue{
		A: []string{},
	}
	if len(defaultValue) > 0 {
		r = &ArrayStringValue{
			A: defaultValue[0],
		}
	}

	switch t := value.(type) {
	case []string:
		r.T = "[]string"
		r.A = t
	case []interface{}:
		r.T = "[]interface {}"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToString(value).A
		}
	case []int:
		r.T = "[]int"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToString(value).A
		}
	case []float64:
		r.T = "[]int"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToString(value).A
		}
	case []int64:
		r.T = "[]int"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToString(value).A
		}
	case []bool:
		r.T = "[]int"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToString(value).A
		}
	case []int8:
		r.T = "[]int"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToString(value).A
		}
	case []int32:
		r.T = "[]int"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToString(value).A
		}
	case []float32:
		r.T = "[]int"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = ToString(value).A
		}
	default:
		r.T = fmt.Sprintf("%T", value)
	}
	return r
}
