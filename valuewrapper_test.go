package hive

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_valueWrapper_singletonValueWrapper(t *testing.T) {
	value := "test"
	wrapper := &singletonValueWrapper{
		fn: func(_ TypeMapper) reflect.Value {
			return reflect.ValueOf(value)
		},
	}

	testValue := wrapper.Get()

	// The type of the value returned should be the original type.
	assert.Equal(t, reflect.String, testValue.Type().Kind())

	// The value should also match.
	assert.Equal(t, value, wrapper.Get().Interface().(string))
}

func Test_valueWrapper_factoryValueWrapper(t *testing.T) {
	value := "test"
	callCount := 0
	wrapper := &factoryValueWrapper{
		fn: func(_ TypeMapper) reflect.Value {
			callCount++
			return reflect.ValueOf(value)
		},
	}

	testFunc := func(callCountShouldBe int) {
		testValue := wrapper.Get()

		assert.Equal(t, reflect.String, testValue.Type().Kind())
		assert.Equal(t, value, testValue.Interface().(string))

		assert.Equal(t, callCountShouldBe, callCount)
	}

	testFunc(1)
	testFunc(2)
	testFunc(3)
}
