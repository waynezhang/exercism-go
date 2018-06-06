package flatten

// Flatten doc here
func Flatten(data interface{}) []interface{} {
	var ret = []interface{}{}

	switch e := data.(type) {
	case []interface{}:
		for _, v := range e {
			ret = append(ret, Flatten(v)...)
		}
	case interface{}:
		ret = append(ret, e)
	}

	return ret
}
