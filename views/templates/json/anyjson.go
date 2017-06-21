package json

type MultiDimension map[string] interface{}
type MapMultiDimension [] map[string] interface{}
type simpleDimension [] interface{}
type stringDimension [] string

func isSimpleDimension(value interface{}) ([] interface{}, bool) {
	switch vv := value.(type) {
	case [] interface{}:
		return vv, true
	}

	return nil, false
}
func isMultiDimension(value interface{}) (map[string] interface{}, bool) {
	switch vv := value.(type) {
	case map[string] interface{}:
		return vv, true
	}

	return nil, false
}
func isMapMultiDimension(value interface{}) ([] map[string] interface{}, bool) {
	switch vv := value.(type) {
	case [] map[string] interface{}:
		return vv, true
	}

	return nil, false
}
