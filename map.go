package typeconv

import (
	"github.com/pkg/errors"
	"reflect"
)

type MapValue struct {
	A         map[string]interface{} // answer
	IsNil     bool
	NotStruct bool
	Error     error
}

func ToMap(value interface{}) (r *MapValue) {
	r = &MapValue{}
	defer func() {
		recv := recover()
		if recv != nil {
			r.Error = errors.Errorf("Error ToMap: %v", recv)
		}
	}()
	if value == nil {
		r.IsNil = true
		return r
	}
	switch t := value.(type) {
	case map[string]interface{}:
		r.A = t
		r.NotStruct = true
		return r
	}
	if reflect.Indirect(reflect.ValueOf(value)).Kind() != reflect.Struct {
		r.NotStruct = true
		return
	}
	r.A = toMap(value)
	return r
}

func toMap(b interface{}) map[string]interface{} {
	t := reflect.TypeOf(b).Elem()
	v := reflect.ValueOf(b).Elem()
	num := t.NumField()
	m := map[string]interface{}{}
	for i := 0; i < num; i++ {
		field := t.Field(i)
		vf := v.Field(i)
		if !vf.CanSet() { // private variable
			continue
		}
		name := field.Name
		vi := vf.Interface()
		if vi == nil {
			continue
		}
		vii := reflect.Indirect(reflect.ValueOf(vi))
		if vii.Kind() != reflect.Struct {
			if !vii.IsValid() { // skip if value is nil
				continue
			}
			m[name] = vi
			continue
		}
		vi = toMap(vi)
		m[name] = vi
	}
	return m
}
