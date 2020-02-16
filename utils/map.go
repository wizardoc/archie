package utils

func MapKeys(m interface{}) interface{} {
	var keys []interface{}

	for k := range m.(map[interface{}]interface{}) {
		keys = append(keys, k)
	}

	return keys
}

func MapValues(m interface{}) interface{} {
	var values []interface{}

	for _, v := range m.(map[interface{}]interface{}) {
		v = append(values, v)
	}

	return values
}
