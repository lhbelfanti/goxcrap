package tweets

// MockRetrieveAll mocks the function MakeRetrieveAll and the values returned by RetrieveAll
func MockRetrieveAll(tweets []Tweet, err error) RetrieveAll {
	return func() ([]Tweet, error) {
		return tweets, err
	}
}

// MockTweet mocks a Tweet
func MockTweet() Tweet {
	return Tweet{
		ID:        "Tweet ID",
		Timestamp: "2024-02-26T18:31:49.000Z",
		IsAReply:  false,
		HasQuote:  false,
		Data: Data{
			HasText:   false,
			HasImages: false,
			Text:      "Tweet Description",
			Images:    []string{"Img 1", "Img 2"},
		},
		Quote: Quote{
			IsAReply: false,
			Data: Data{
				HasText:   false,
				HasImages: false,
				Text:      "",
				Images:    nil,
			},
		},
	}
}
