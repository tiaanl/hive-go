package hive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type anInterface interface{}

type notAnInterface struct{}

func Test_InterfaceOf_PanicIfWePassIncorrectType(t *testing.T) {
	// Panic if we pass a struct.
	assert.Panics(t, func() {
		InterfaceOf(notAnInterface{})
	})

	// Panic if we pass a pointer to a struct.
	assert.Panics(t, func() {
		InterfaceOf(&notAnInterface{})
	})

	// Panic if it's an interface, but not a pointer.
	assert.Panics(t, func() {
		InterfaceOf((anInterface)(nil))
	})

	// Don't panic if we pass a double pointer.
	assert.Equal(t, "anInterface", InterfaceOf((**anInterface)(nil)).Name())
}
