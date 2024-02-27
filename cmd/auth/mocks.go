package auth

// MockLogin mocks the function MakeLogin and the values returned by Login
func MockLogin(err error) Login {
	return func() error {
		return err
	}
}
