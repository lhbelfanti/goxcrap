package corpuscreator

import "context"

// MockSaveTweets mocks SaveTweets function
func MockSaveTweets(err error) SaveTweets {
	return func(ctx context.Context, body SaveTweetsBody) error {
		return err
	}
}

// MockGetSearchCriteriaExecution mocks GetSearchCriteriaExecution function
func MockGetSearchCriteriaExecution(execution Execution, err error) GetSearchCriteriaExecution {
	return func(ctx context.Context, executionID int) (Execution, error) {
		return execution, err
	}
}

// MockUpdateSearchCriteriaExecution mocks UpdateSearchCriteriaExecution function
func MockUpdateSearchCriteriaExecution(err error) UpdateSearchCriteriaExecution {
	return func(ctx context.Context, executionID int, body UpdateSearchCriteriaExecutionBody) error {
		return err
	}
}

// MockInsertSearchCriteriaExecutionDay mocks InsertSearchCriteriaExecutionDay function
func MockInsertSearchCriteriaExecutionDay(err error) InsertSearchCriteriaExecutionDay {
	return func(ctx context.Context, executionID int, body InsertSearchCriteriaExecutionDayBody) error {
		return err
	}
}

// MockSaveTweetsBody mocks a SaveTweetsBody
func MockSaveTweetsBody() SaveTweetsBody {
	return SaveTweetsBody{
		MockTweetDTO(),
		MockTweetDTO(),
	}
}

// MockTweetDTO mocks a TweetDTO
func MockTweetDTO() TweetDTO {
	textContent := "test"
	quote := MockQuoteDTO()
	searchCriteriaID := 1

	return TweetDTO{
		StatusID:         "123456789012345",
		IsAReply:         true,
		TextContent:      &textContent,
		Images:           []string{"test1", "test2"},
		Quote:            &quote,
		SearchCriteriaID: &searchCriteriaID,
	}
}

// MockQuoteDTO mocks a QuoteDTO
func MockQuoteDTO() QuoteDTO {
	textContent := "test"

	return QuoteDTO{
		IsAReply:    true,
		TextContent: &textContent,
		Images:      []string{"test1", "test2"},
	}
}

// MockExecution mocks an Execution
func MockExecution(status string) Execution {
	return Execution{
		ID:               1,
		Status:           status,
		SearchCriteriaID: 2,
	}
}

// MockUpdateSearchCriteriaExecutionBody mocks a UpdateSearchCriteriaExecutionBody
func MockUpdateSearchCriteriaExecutionBody() UpdateSearchCriteriaExecutionBody {
	return UpdateSearchCriteriaExecutionBody{
		Status: "DONE",
	}
}

// MockInsertSearchCriteriaExecutionDayBody mocks a InsertSearchCriteriaExecutionDayBody
func MockInsertSearchCriteriaExecutionDayBody() InsertSearchCriteriaExecutionDayBody {
	return InsertSearchCriteriaExecutionDayBody{
		ExecutionDate:             "01/01/2024",
		TweetsQuantity:            10,
		ErrorReason:               nil,
		SearchCriteriaExecutionID: 1,
	}
}
