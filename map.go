package typeconv

func ToMap(value interface{}) map[string]interface{} {
	switch t := value.(type) {
	case map[string]interface{}:
		return t
	}
	return nil
}
