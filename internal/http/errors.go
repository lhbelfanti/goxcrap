package http

import "errors"

var (
	FailedToMarshalBody    = errors.New("failed to marshal body")
	FailedToCreateRequest  = errors.New("failed to create request")
	FailedToExecuteRequest = errors.New("failed to execute request")
	FailedToReadResponse   = errors.New("failed to read response")
)
