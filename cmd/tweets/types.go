package tweets

type (
	// Tweet contains the information needed to represent a tweet
	Tweet struct {
		ID        string
		Timestamp string
		IsAReply  bool
		HasQuote  bool
		Data
		Quote
	}

	// Data defines the selected parts of a tweets that will be saved
	Data struct {
		HasText   bool
		HasImages bool
		Text      string
		Images    []string
	}

	// Quote contains the information of a retweeted tweet in the original tweet
	Quote struct {
		IsAReply bool
		Data
	}
)
