package corpuscreator

import "errors"

var (
	FailedToExecuteRequest    = errors.New("request failed")
	FailedToUnmarshalResponse = errors.New("failed to unmarshal response")
)
