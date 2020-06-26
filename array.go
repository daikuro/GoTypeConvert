package typeconv

type ArrayValue struct {
	A []interface{} // value
}

func ToArray(value interface{}, defaultValue ...[]interface{}) *ArrayValue {
	r := &ArrayValue{
		A: []interface{}{},
	}
	if len(defaultValue) > 0 {
		r = &ArrayValue{
			A: defaultValue[0],
		}
	}

	switch t := value.(type) {
	case []interface{}:
		r.A = t
	}
	return r
}
