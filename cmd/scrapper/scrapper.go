package scrapper

import (
	"fmt"

	"goxcrap/cmd/auth"
	"goxcrap/cmd/search"
)

// Execute starts the twitter scrapper
func Execute(login auth.Login, getAdvanceSearchCriteria search.GetAdvanceSearchCriteria, executeAdvanceSearch search.ExecuteAdvanceSearch) error {
	err := login()
	if err != nil {
		return err
	}
	fmt.Println("Log In completed")

	searchCriteria := getAdvanceSearchCriteria()
	for _, criteria := range searchCriteria {
		fmt.Printf("Criteria: '%s'\n", criteria.ID)
		since, until, err := criteria.ParseDates()
		if err != nil {
			fmt.Printf("Error parsing dates: Since %s - Until %s - Error %v\n", criteria.Since, criteria.Until, err)
			continue
		}

		currentCriteria := criteria
		for current := since; !current.After(until); current = current.AddDays(1) {
			currentCriteria.Since = current.String()
			currentCriteria.Until = current.AddDays(1).String()
			err := executeAdvanceSearch(currentCriteria)
			if err != nil {
				fmt.Printf("Error while executing advance search - Error %v\n", err)
				continue
			}

			// TODO: retrieve tweets of the current page
			// TODO: obtain tweets information (text, images, links)
			// TODO: save tweets
		}

		fmt.Printf("All the tweets of the criteria '%s' were retrieved", criteria.ID)
	}

	return nil
}
