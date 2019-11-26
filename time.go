package typeconv

import "time"

type TimeValue struct {
	A       time.Time // answer
	Error   error
	NoValue bool
}

func ToTime(v interface{}, defaultValue ...time.Time) *TimeValue {
	r := &TimeValue{
		NoValue: false,
	}
	switch t := v.(type) {
	case time.Time:
		r.A = t
	default:
		if len(defaultValue) > 0 {
			r.A = defaultValue[0]
		}
		r.NoValue = true
	}

	return r
}
