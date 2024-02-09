package setup

// Must function throws a panic if the initialization of a dependency fails
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}

	return val
}
