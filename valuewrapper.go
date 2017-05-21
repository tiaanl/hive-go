package hive

import "reflect"

type valueWrapper interface {
	Get() reflect.Value
}

type instanceValueWrapper struct {
	value reflect.Value
}

func (ivw *instanceValueWrapper) Get() reflect.Value {
	return ivw.value
}

type funcValueWrapper struct {
	container Container
	fn        func(TypeMapper) interface{}
}

func (fvw *funcValueWrapper) Get() reflect.Value {
	return reflect.ValueOf(fvw.fn(fvw.container))
}
