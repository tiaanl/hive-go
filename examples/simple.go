package main

import (
	"fmt"
	"reflect"

	"github.com/tiaanl/hive-go"
)

type Repository interface {
	Get() string
}

type repository struct{}

func (*repository) Get() string {
	return "test"
}

func main() {
	container := hive.New()

	// The user repository.
	usersRepository := &repository{}

	// Set the users repository in the container.
	container.Set(reflect.TypeOf(usersRepository), reflect.ValueOf(usersRepository))

	// Now we can get the repository from the container by it's interface.
	resultValue := container.Get(hive.InterfaceOf((*Repository)(nil)))

	// We can cast the value back to the real pointer.
	realUsersRepository := resultValue.Interface().(Repository)

	fmt.Println(realUsersRepository.Get())
}
