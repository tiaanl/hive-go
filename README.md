# hive-go
Dependency Injection container for Go

## Example

```go
package main

import (
    "reflect"
    "github.com/tiaanl/hive-go"
)

func main() {
    // Store a users repository in the container.
    container := hive.New()
    usersRepository := users.NewRepository()
    container.Set(hive.InterfaceOf((*users.Repository)(nil)), reflect.ValueOf(usersRepository))
}
```
