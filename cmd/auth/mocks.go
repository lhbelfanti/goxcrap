package auth

// MockLogin mocks Login function
func MockLogin(err error) Login {
	return func() error {
		return err
	}
}
