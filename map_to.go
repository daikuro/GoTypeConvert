package typeconv

import "reflect"

func MapToInterface(o interface{}, m map[string]interface{}) interface{} {
	t := reflect.TypeOf(o).Elem()
	p := reflect.New(t)
	v := reflect.Indirect(p)
	num := t.NumField()
	for i := 0; i < num; i++ {
		field := t.Field(i)
		value := v.Field(i)
		name := field.Name
		if m[name] == nil {
			continue
		}
		value.Set(reflect.ValueOf(m[name]))
	}
	return p.Interface()
}
