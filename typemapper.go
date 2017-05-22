package hive

import "reflect"

// ValueFunc represents a function that return a reflect.Value of the object that you want to store
// in a container.
type ValueFunc func(TypeMapper) reflect.Value

// TypeMapper is the interface for storing and retrieving values from a container.
type TypeMapper interface {
	// GetString the value stored in the container based on the type provided.
	Get(reflect.Type) reflect.Value

	// Set a value in the container that acts as a singleton.  Getting the value multiple times will
	// always return the same instance.
	Singleton(reflect.Type, ValueFunc) TypeMapper
}
