package tweets

import "github.com/tebeka/selenium"

// MockRetrieveAll mocks RetrieveAll function
func MockRetrieveAll(tweets []Tweet, err error) RetrieveAll {
	return func() ([]Tweet, error) {
		return tweets, err
	}
}

// MockTweet mocks a Tweet
func MockTweet() Tweet {
	return Tweet{
		ID:        "02bd92faa38aaa6cc0ea75e59937a1ef8d6ad3a9f75f3ac4166fef23da9f209b",
		Timestamp: "2024-02-26T18:31:49.000Z",
		IsAReply:  true,
		HasQuote:  true,
		Data: Data{
			HasText:   true,
			HasImages: true,
			Text:      "Tweet Description",
			Images:    []string{"Img 1", "Img 2"},
		},
		Quote: Quote{
			IsAReply: true,
			Data: Data{
				HasText:   true,
				HasImages: true,
				Text:      "Quote Description",
				Images:    []string{"Img 3", "Img 4"},
			},
		},
	}
}

// MockGetAuthor mocks GetAuthor function
func MockGetAuthor(author string, err error) GetAuthor {
	return func(element selenium.WebElement) (string, error) {
		return author, err
	}
}

// MockGetTimestamp mocks GetTimestamp function
func MockGetTimestamp(timestamp string, err error) GetTimestamp {
	return func(element selenium.WebElement) (string, error) {
		return timestamp, err
	}
}
