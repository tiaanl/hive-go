package hive

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TypeMapper_SettingAndGetting(t *testing.T) {
	type Object interface{}
	type object struct {
		a int
	}

	value := &object{a: 10}

	container := New()
	container.Singleton(InterfaceOf((*Object)(nil)), func(_ TypeMapper) reflect.Value {
		return reflect.ValueOf(value)
	})

	testValue := container.Get(InterfaceOf((*Object)(nil)))
	assert.True(t, testValue.IsValid())
	assert.False(t, testValue.IsNil())
	assert.Equal(t, 10, testValue.Interface().(*object).a)
}

func Test_TypeMapper_MakeSureLazyIsLazy(t *testing.T) {
	type Object interface{}
	type object struct {
		a int
	}

	createCount := 0

	container := New()
	container.Singleton(InterfaceOf((*Object)(nil)), func(c TypeMapper) reflect.Value {
		newValue := &object{a: 10}
		createCount = createCount + 1
		return reflect.ValueOf(newValue)
	})

	// At this point we should not have created a new object yet.
	assert.Equal(t, 0, createCount)

	testValue := container.Get(InterfaceOf((*Object)(nil)))
	original := testValue.Interface().(*object)

	// Make sure we have the original object back.
	assert.Equal(t, 10, original.a)

	// Make sure it was only created once.
	assert.Equal(t, 1, createCount)

	// When we get the interface a second time, we should not create it again.
	testValue = container.Get(InterfaceOf((*Object)(nil)))
	assert.Equal(t, 1, createCount)
}

// The interface we'll use as the key.
type stringGetter interface {
	GetString() string
}

type intGetter interface {
	GetInt() int
}

// We create a struct that will implement the above interface.
type realGetter struct {
	stringValue string
	intValue    int
}

func (r *realGetter) GetString() string {
	return r.stringValue
}

func (r *realGetter) GetInt() int {
	return r.intValue
}

func Test_TypeMapper_GetValueByInterfaceImplemented(t *testing.T) {
	container := New()

	value := &realGetter{
		stringValue: "real",
		intValue:    10,
	}

	// Store the value in the container using it's own type.
	container.Singleton(reflect.TypeOf(value), func(_ TypeMapper) reflect.Value {
		return reflect.ValueOf(value)
	})

	// See if we can get it as a stringGetter.
	testValue := container.Get(InterfaceOf((*stringGetter)(nil)))
	assert.Equal(t, value.stringValue, testValue.Interface().(stringGetter).GetString())

	// Also test if we can get it as an intGetter.
	testValue = container.Get(InterfaceOf((*intGetter)(nil)))
	assert.Equal(t, value.intValue, testValue.Interface().(intGetter).GetInt())
}

func Test_TypeMapper_GetValueFromParent(t *testing.T) {
	parentContainer := New()

	value := &realGetter{
		stringValue: "real",
		intValue:    10,
	}

	parentContainer.Singleton(reflect.TypeOf(value), func(_ TypeMapper) reflect.Value {
		return reflect.ValueOf(value)
	})

	// Create a child container, with the original as the parent.
	childContainer := NewWithParent(parentContainer)

	// Now get the realGetter from the childContainer and we should get the one
	// from the parent.
	testValue := childContainer.Get(reflect.TypeOf(value))
	assert.Equal(t, value.stringValue, testValue.Interface().(stringGetter).GetString())
	assert.Equal(t, value.intValue, testValue.Interface().(intGetter).GetInt())
}

func Test_TypeMapper_GetChildValueIfItExistsInTheChildAndParent(t *testing.T) {
	// Create a parent container that will have a realGetter inside with it's
	// real type as the key.
	parentContainer := New()
	parentValue := &realGetter{
		stringValue: "parent",
	}
	parentContainer.Singleton(reflect.TypeOf(parentValue), func(_ TypeMapper) reflect.Value {
		return reflect.ValueOf(parentValue)
	})

	// Create a child container and set a new realGetter in it with the same
	// type as the one in the parent.
	childContainer := NewWithParent(parentContainer)
	childValue := &realGetter{
		stringValue: "child",
	}
	childContainer.Singleton(reflect.TypeOf(childValue), func(_ TypeMapper) reflect.Value {
		return reflect.ValueOf(childValue)
	})

	// If we get the value from the container, then we should get the child value.
	testValue := childContainer.Get(reflect.TypeOf(childValue))
	assert.Equal(t, childValue.stringValue, testValue.Interface().(stringGetter).GetString())
}

func Test_TypeMapper_DoNotCrashIfTheTypeIsNotFound(t *testing.T) {
	container := New()

	testValue := container.Get(InterfaceOf((*stringGetter)(nil)))

	assert.False(t, testValue.IsValid())
}
