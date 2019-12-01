package typeconv

import (
	"github.com/pkg/errors"
	"reflect"
)

type MapToValue struct {
	A     interface{} // answer
	IsNil bool
	Error error
}

func MapToInterface(p interface{}, m map[string]interface{}) (r *MapToValue) {
	r = &MapToValue{}
	defer func() {
		recv := recover()
		if recv != nil {
			r.Error = errors.Errorf("Error MapToInterface: %v", recv)
		}
	}()
	if m == nil {
		r.IsNil = true
		return
	}
	t := reflect.TypeOf(p).Elem()
	v := reflect.Indirect(reflect.ValueOf(p))
	num := t.NumField()
	for i := 0; i < num; i++ {
		newValue := m[t.Field(i).Name]
		if newValue == nil {
			continue
		}
		v.Field(i).Set(reflect.ValueOf(newValue))
	}
	r.A = p
	return
}
