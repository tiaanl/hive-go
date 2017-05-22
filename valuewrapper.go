package hive

import "reflect"

type valueWrapper interface {
	Get() reflect.Value
}

type singletonValueWrapper struct {
	value reflect.Value
}

func (w *singletonValueWrapper) Get() reflect.Value {
	return w.value
}

type lazySingletonValueWrapper struct {
	container Container
	fn        func(TypeMapper) interface{}
	value     reflect.Value
}

func (w *lazySingletonValueWrapper) Get() reflect.Value {
	// Cache the value so that we are a singleton.
	if w.value.IsValid() {
		return w.value
	}
	w.value = reflect.ValueOf(w.fn(w.container))
	return w.value
}
