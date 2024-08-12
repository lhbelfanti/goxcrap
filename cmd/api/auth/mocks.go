package auth

import "context"

// MockLogin mocks Login function
func MockLogin(err error) Login {
	return func(ctx context.Context) error {
		return err
	}
}
