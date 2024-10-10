package scrapper

import "errors"

const (
	FailedToRunScrapper string = "Failed to run scrapper"

	CantDecodeBodyIntoCriteria string = "Can't decode body into a criteria"
	CantReEnqueueFailedMessage string = "Can't re-enqueue failed message"
)

var (
	FailedToLogin                          = errors.New("failed to login")
	FailedToUpdateSearchCriteriaExecution  = errors.New("failed to update search criteria execution")
	FailedToParseDatesFromTheGivenCriteria = errors.New("failed to parse dates from the given criteria")

	FailedToDecodeBodyIntoCriteria              = errors.New("failed to decode body into a criteria")
	FailedToRetrieveSearchCriteriaExecutionData = errors.New("failed to retrieve search criteria execution data")
	FailedToReEnqueueFailedMessage              = errors.New("failed to re-enqueue failed message")
	FailedToRunScrapperProcess                  = errors.New("failed to run scrapper")
)
