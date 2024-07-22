package scrapper

import "errors"

var (
	FailedToLogin                          = errors.New("failed to login")
	FailedToParseDatesFromTheGivenCriteria = errors.New("failed to parse dates from the given criteria")
)

const (
	FailedToRunScrapper string = "Failed to run scrapper"

	CantDecodeBodyIntoCriteria string = "Can't decode body into a criteria"
	CantReEnqueueFailedMessage string = "Can't re-enqueue failed message"
)
