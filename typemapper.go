package hive

import "reflect"

type TypeMapper interface {
	// Set a value in the container linked to the type provided.  When getting
	// this value from the container, the caller will always receive the same
	// instance of the value.
	Singleton(reflect.Type, reflect.Value) TypeMapper

	// Set a factory function for the provided type.  This will allow the
	// instance only to be created once the type is requested.
	LazySingleton(reflect.Type, func(TypeMapper) interface{}) TypeMapper

	// GetString the value stored in the container based on the type provided.
	Get(reflect.Type) reflect.Value
}
