package hive

var (
	defaultContainer Container
)

func DefaultContainer() Container {
	if defaultContainer == nil {
		defaultContainer = New()
	}

	return defaultContainer
}
