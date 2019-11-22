package typeconv

import (
	"reflect"
)

func ToMap(value interface{}) map[string]interface{} {
	if value == nil {
		return nil
	}
	switch t := value.(type) {
	case map[string]interface{}:
		return t
	}
	if reflect.Indirect(reflect.ValueOf(value)).Kind() != reflect.Struct {
		return nil
	}
	return toMap(value)
}

func toMap(b interface{}) map[string]interface{} {
	if b == nil {
		return nil
	}
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
