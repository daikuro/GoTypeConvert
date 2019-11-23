package typeconv

import "reflect"

func MapToInterface(p interface{}, m map[string]interface{}) interface{} {
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
	return p
}
