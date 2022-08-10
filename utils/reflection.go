package utils

import (
	"reflect"
)

func Implements(e interface{}, typeInterface reflect.Type) bool {
	return reflect.TypeOf(e).Implements(typeInterface)
}
