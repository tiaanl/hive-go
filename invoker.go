package hive

import "reflect"

type Invoker interface {
	Invoke(interface{}) ([]reflect.Value, error)
}
