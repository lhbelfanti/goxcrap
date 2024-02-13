package scrapper

import (
	"fmt"

	"goxcrap/cmd/auth"
)

// Init starts the twitter scrapper
func Init(login auth.Login) error {
	err := login()
	if err != nil {
		return err
	}

	fmt.Println("Login completed")

	return nil
}
