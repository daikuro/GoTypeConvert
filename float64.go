package typeconv

import "strconv"

type Float64Value struct {
	A     float64 // answer
	Error error
}

var DefaultFloat64 = 0.0

func ToFloat64(value interface{}) *Float64Value {
	return ToFloat64d(value, DefaultFloat64)
}

func ToFloat64d(value interface{}, defaultValue float64) *Float64Value {
	r := &Float64Value{
		A: defaultValue,
	}
	switch t := value.(type) {
	case float32:
		r.A = float64(t)
	case float64:
		r.A = t
	case string:
		f, e := strconv.ParseFloat(t, 64)
		r.A = f
		r.Error = e
	case int:
		r.A = float64(t)
	case int8:
		r.A = float64(t)
	case int16:
		r.A = float64(t)
	case int32:
		r.A = float64(t)
	case int64:
		r.A = float64(t)
	}

	return r
}
