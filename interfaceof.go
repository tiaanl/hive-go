package hive

import "reflect"

// To get a reflect.Type object pointing to an interface, pass in a pointer to an interface typed
// object. Example:
//
// InterfaceOf((*Repository)(nil))
//
// will return a reflect.Type object with Kind() of Repository.
func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("InterfaceOf called with a value that is not a pointer to an interface.")
	}

	return t
}
