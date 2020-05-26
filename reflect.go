package gen

import "reflect"

func IsZero(i interface{}) bool {
	v := reflect.ValueOf(i)
	return !v.IsValid() || reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
