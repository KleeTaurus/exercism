package flatten

import (
	"reflect"
)

func Flatten(nested interface{}) []interface{} {
	flatten := make([]interface{}, 0)

	v := reflect.ValueOf(nested)
	switch v.Kind() {
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			flatten = append(flatten, Flatten(v.Index(i).Interface())...)
		}
	case reflect.Int:
		flatten = append(flatten, v.Interface())
	default:
		// just ignore everything else
	}

	return flatten
}
