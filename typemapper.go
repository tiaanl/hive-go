package hive

import "reflect"

type TypeMapper interface {
	// Set the value for the provided type.
	Set(reflect.Type, reflect.Value) TypeMapper

	// Set a factory function for the provided type.  This will allow the instance only to be created once the type
	// is requested.
	SetLazy(reflect.Type, func(TypeMapper) interface{}) TypeMapper

	// GetString the value stored in the container based on the type provided.
	Get(reflect.Type) reflect.Value
}
