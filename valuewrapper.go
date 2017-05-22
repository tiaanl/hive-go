package hive

import "reflect"

type valueWrapper interface {
	Get() reflect.Value
}

// This wrapper caches the result it gets from Get so that it always returns the same instance.
type singletonValueWrapper struct {
	container Container
	fn        ValueFunc
	value     reflect.Value
}

func (w *singletonValueWrapper) Get() reflect.Value {
	// Cache the value so that we are a singleton.
	if w.value.IsValid() {
		return w.value
	}

	w.value = w.fn(w.container)
	return w.value
}

// This wrapper always calls the function when getting the value.
type factoryValueWrapper struct {
	container Container
	fn        ValueFunc
}

func (w *factoryValueWrapper) Get() reflect.Value {
	return w.fn(w.container)
}
