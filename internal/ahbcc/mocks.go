package ahbcc

// MockSaveTweets mocks SaveTweets function
func MockSaveTweets(err error) SaveTweets {
	return func(body SaveTweetsBody) error {
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
	hash := "02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b"
	textContent := "test"
	quote := MockQuoteDTO()
	searchCriteriaID := 1

	return TweetDTO{
		Hash:             &hash,
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
