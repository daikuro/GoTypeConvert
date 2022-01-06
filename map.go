package typeconv

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"time"
)

type MapValue struct {
	A         map[string]interface{} // answer
	IsNil     bool
	NotStruct bool
	Error     error
}

type SetMapDirectFunc = func(name string, v interface{}) bool

var DefaultMapSet = func(name string, v interface{}) bool {
	switch v.(type) {
	case time.Time, *time.Time:
		return true
	}
	return false
}

func ToMap(value interface{}) (r *MapValue) {
	return ToTagMap(value, "")
}

func ToTagMap(value interface{}, tag string) (r *MapValue) {
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
	r.A = toMap(value, DefaultMapSet, tag)
	return r
}

func toMap(b interface{}, defaultSet SetMapDirectFunc, tag string) map[string]interface{} {
	m := map[string]interface{}{}
	t := reflect.TypeOf(b)
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		t = t.Elem()
	}
	v := reflect.ValueOf(b)
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		v = v.Elem()
	}
	num := t.NumField()
	for i := 0; i < num; i++ {
		field := t.Field(i)
		name := field.Name
		if tag != "" {
			v, _ := field.Tag.Lookup(tag)
			if v != "" {
				ss := strings.Split(v, ",")
				name = ss[0]
				if name == "-" {
					continue
				}
			}
		}
		vf := v.Field(i)
		if !vf.CanSet() { // private variable
			continue
		}
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
		if defaultSet(name, vi) {
			m[name] = vi
			continue
		}
		vi = toMap(vi, defaultSet, tag)
		m[name] = vi
	}
	return m
}
