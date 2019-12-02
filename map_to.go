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

var ArrayStringType = reflect.TypeOf([]string{})
var ArrayInterfaceType = reflect.TypeOf([]interface{}{})

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
		switch newValue.(type) {
		case []interface{}:
			switch v.Field(i).Type() {
			case ArrayStringType:
				v.Field(i).Set(reflect.ValueOf(ToArrayString(newValue).A))
			case ArrayInterfaceType:
				v.Field(i).Set(reflect.ValueOf(newValue))
			}

		default:
			v.Field(i).Set(reflect.ValueOf(newValue))
		}
	}
	r.A = p
	return
}
