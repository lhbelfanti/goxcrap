package scrapper

import "errors"

var (
	FailedToParseDatesFromTheGivenCriteria = errors.New("failed to parse dates from the given criteria")
)

const (
	FailedToRunScrapper string = "Failed to run scrapper"
	InvalidBody         string = "Can't parse criteria from body"
)
