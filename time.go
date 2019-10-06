package typeconv

import "time"

type TimeValue struct {
	A       time.Time // answer
	Error   error
	NoValue bool
}

func ToTime(v interface{}) *TimeValue {
	r := &TimeValue{
		NoValue: false,
	}
	switch t := v.(type) {
	case time.Time:
		r.A = t
	default:
		r.NoValue = true
	}

	return r
}
