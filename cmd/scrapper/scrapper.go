package scrapper

import (
	"fmt"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/search"
)

// Execute starts the twitter scrapper
func Execute(login auth.Login, getSearchCriteria search.GetSearchCriteria) error {
	err := login()
	if err != nil {
		return err
	}
	fmt.Println("Log In completed")

	searchCriteria := getSearchCriteria()
	for _, criteria := range searchCriteria {
		fmt.Printf("Search criteria: '%s'", criteria.ID)

		// TODO: search
		// TODO: save tweet
		// fmt.Printf("All the tweets of the criteria '%s' were retrieved", criteria.ID)
	}

	return nil
}
