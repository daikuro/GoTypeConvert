package typeconv

import "fmt"

var DefaultString = ""

func ToString(value interface{}) string {
	return ToStringd(value, DefaultString)
}

func ToStringd(value interface{}, defaultValue string) string {
	if value == nil {
		return defaultValue
	}
	switch t := value.(type) {
	case string:
		return t
	default:
		return fmt.Sprint(value)
	}
}
