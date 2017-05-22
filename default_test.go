package hive

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultContainer_MultipleCallsShouldReturnTheSameContainer(t *testing.T) {
	container := DefaultContainer()

	value := "test"

	container.Singleton(reflect.TypeOf(value), reflect.ValueOf(value))

	func() {
		container := DefaultContainer()

		testValue := container.Get(reflect.TypeOf(value))
		assert.Equal(t, value, testValue.Interface().(string))
	}()
}
