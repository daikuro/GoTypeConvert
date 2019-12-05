package typeconv

import (
	"fmt"
	"strconv"
)

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
			r.A[i] = strconv.Itoa(value)
		}
	case []int8:
		r.T = "[]int8"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = strconv.FormatInt(int64(value), 10)
		}
	case []int32:
		r.T = "[]int32"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = strconv.FormatInt(int64(value), 10)
		}
	case []int64:
		r.T = "[]int64"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = strconv.FormatInt(value, 10)
		}
	case []float32:
		r.T = "[]float32"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = strconv.FormatFloat(float64(value), 'G', 4, 32)
		}
	case []float64:
		r.T = "[]float64"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = strconv.FormatFloat(value, 'G', 4, 64)
		}
	case []bool:
		r.T = "[]bool"
		r.A = make([]string, len(t))
		for i, value := range t {
			r.A[i] = strconv.FormatBool(value)
		}
	default:
		r.T = fmt.Sprintf("%T", value)
	}
	return r
}
