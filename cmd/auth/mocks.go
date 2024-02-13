package auth

// MockMakeLogin mocks the function MakeLogin and the values returned by Login
func MockMakeLogin(err error) Login {
	return func() error {
		return err
	}
}
