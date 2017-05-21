package hive

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Invoker_WeGetTheValuesWeAskFor(t *testing.T) {
	stringValue := "test"
	intValue := 10

	container := New()
	container.Set(reflect.TypeOf(stringValue), reflect.ValueOf(stringValue))
	container.Set(reflect.TypeOf(intValue), reflect.ValueOf(intValue))

	testValue := ""

	_, err := container.Invoke(func(s string, i int) {
		testValue = fmt.Sprintf("%s%d", s, i)
	})
	assert.NoError(t, err)
	assert.Equal(t, "test10", testValue)
}

func Test_Invoker_WeGetReturnValues(t *testing.T) {
	value := "test"

	container := New()
	container.Set(reflect.TypeOf(value), reflect.ValueOf(value))

	returnValues, err := container.Invoke(func(s string) string {
		return s + s
	})

	assert.NoError(t, err)
	assert.Equal(t, 1, len(returnValues))
	assert.Equal(t, "testtest", returnValues[0].Interface().(string))
}

func Test_Invoker_GetErrorWhenTypeNotFound(t *testing.T) {
	container := New()
	returnValues, err := container.Invoke(func(i int) int {
		return i
	})

	assert.Nil(t, returnValues)
	assert.Error(t, err)
}
