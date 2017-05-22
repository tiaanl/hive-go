package hive

import (
	"fmt"
	"reflect"
)

type Container interface {
	TypeMapper
	Invoker
}

func New() Container {
	return &container{
		values: make(map[reflect.Type]valueWrapper),
	}
}

func NewWithParent(parent Container) Container {
	return &container{
		values: make(map[reflect.Type]valueWrapper),
		parent: parent,
	}
}

type container struct {
	values map[reflect.Type]valueWrapper
	parent Container
}

func (c *container) Singleton(t reflect.Type, v reflect.Value) TypeMapper {
	c.values[t] = &singletonValueWrapper{
		value: v,
	}
	return c
}

func (c *container) LazySingleton(t reflect.Type, fn func(TypeMapper) interface{}) TypeMapper {
	c.values[t] = &lazySingletonValueWrapper{
		container: c,
		fn:        fn,
	}
	return c
}

func (c *container) Get(t reflect.Type) reflect.Value {
	// Get a value from the map that has the exact type provided.
	wrapper, ok := c.values[t]

	// If we can't find an exact match and we asked for an interface, then look
	// in the container for a value that implements the interface we're looking
	// for.
	if !ok && t.Kind() == reflect.Interface {
		for k, v := range c.values {
			if k.Implements(t) {
				wrapper = v
				break
			}
		}
	}

	// If we found a wrapper, we just return it's value.
	if wrapper != nil {
		return wrapper.Get()
	}

	// We could not satisfy the request for a value, but if we have a parent,
	// then we check that first.
	if c.parent != nil {
		return c.parent.Get(t)
	}

	// Oops, the type wasn't found.
	return reflect.ValueOf(nil)
}

func (c *container) Invoke(f interface{}) ([]reflect.Value, error) {
	t := reflect.TypeOf(f)

	var in = make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		argType := t.In(i)
		val := c.Get(argType)
		if !val.IsValid() {
			return nil, fmt.Errorf("Value not found for type %v", argType)
		}

		in[i] = val
	}

	return reflect.ValueOf(f).Call(in), nil
}
