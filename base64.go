package typeconv

import "encoding/base64"

type Base64Value struct {
	A     string // answer
	IsNil bool
}

func ToBase64(v interface{}) (r *Base64Value) {
	r = &Base64Value{}
	if v == nil {
		r.IsNil = true
		return
	}
	switch t := v.(type) {
	case []byte:
		r.A = base64.StdEncoding.EncodeToString(t)
	case string:
		r.A = base64.StdEncoding.EncodeToString([]byte(t))
	}
	return
}

type Base64ToValue struct {
	A     []byte // answer
	Error error
	IsNil bool
}

func Base64To(v interface{}) (r *Base64ToValue) {
	r = &Base64ToValue{}
	if v == nil {
		r.IsNil = true
		return
	}
	switch t := v.(type) {
	case string:
		r.A, r.Error = base64.StdEncoding.DecodeString(t)
		return
	}
	return
}
