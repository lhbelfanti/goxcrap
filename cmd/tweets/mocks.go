package tweets

// MockMakeRetrieveAll mocks the function MakeRetrieveAll and the values returned by RetrieveAll
func MockMakeRetrieveAll(tweets []Tweet, err error) RetrieveAll {
	return func() ([]Tweet, error) {
		return tweets, err
	}
}

// MockTweet mocks a Tweet
func MockTweet() Tweet {
	return Tweet{
		ID:     "Tweet ID",
		Text:   "Tweet Description",
		Images: []string{"Img 1", "Img 2"},
	}
}
