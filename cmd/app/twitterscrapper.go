package app

import (
	"fmt"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/scrapper"
)

// Init starts the twitter scrapper
func Init(s scrapper.Scrapper) error {
	err := auth.Login(s)
	if err != nil {
		return err
	}

	fmt.Println("Login completed")

	return nil
}
