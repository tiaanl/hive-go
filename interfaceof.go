package hive

import "reflect"

func InterfaceOf(value interface{}) reflect.Type {
	t := reflect.TypeOf(value)

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Interface {
		panic("InterfaceOf called with a value that is not a pointer to an interface")
	}

	return t
}
