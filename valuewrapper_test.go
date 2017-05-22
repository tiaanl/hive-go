package hive

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_valueWrapper_funcValueWrapper(t *testing.T) {
	var value string = "test"
	wrapper := &lazySingletonValueWrapper{
		container: nil,
		fn: func(container TypeMapper) interface{} {
			return value
		},
	}

	assert.Equal(t, value, wrapper.Get().Interface().(string))
}

func Test_valueWrapper_singletonValueWrapper(t *testing.T) {
	var value string = "test"
	wrapper := &singletonValueWrapper{
		value: reflect.ValueOf(value),
	}

	assert.Equal(t, value, wrapper.Get().Interface().(string))
}
