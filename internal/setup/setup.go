package setup

// Init function throws a panic if the initialization of a dependency fails
func Init[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}

// Must function throws a panic if the initialization of a dependency fails
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
