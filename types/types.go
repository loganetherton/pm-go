package types

import "reflect"

/*
*
Interface created from a type which can be used to check if another interface implements this one
*/
var nilErr = (*error)(nil)
var ErrorInterface = reflect.TypeOf(nilErr).Elem()

var StringInterface = reflect.TypeOf((*string)(nil)).Elem()
