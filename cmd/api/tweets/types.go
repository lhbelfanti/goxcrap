package tweets

import "fmt"

type (
	// Tweet contains the information needed to represent a tweet
	Tweet struct {
		ID       string
		HasQuote bool
		Data
		Quote
	}

	// Data defines the selected parts of a tweets that will be saved
	Data struct {
		Author    string
		Avatar    string
		Timestamp string
		IsAReply  bool
		HasText   bool
		HasImages bool
		Text      string
		Images    []string
	}

	// Quote contains the information of a retweeted tweet in the original tweet
	Quote struct {
		Data
	}

	// TweetHash contains the ID (hash) calculated with the author and the Timestamp (which is also contained in this struct)
	TweetHash struct {
		ID        string
		Timestamp string
	}
)

// String converts Tweet properties to a string
func (tweet Tweet) String() string {
	return fmt.Sprintf("\n------------------------\n--- Tweet ---\n ID: %s \n HasQuote: %t %s %s\n------------------------\n\n",
		tweet.ID, tweet.HasQuote, tweet.Data.String(), tweet.Quote.String())
}

// String converts Data properties to a string
func (data Data) String() string {
	return fmt.Sprintf("\n   --- Data ---\n   "+
		"Author: %s \n "+
		"Avatar: %s \n "+
		"Timestamp: %s \n "+
		"IsAReply: %t \n "+
		"HasText: %t \n   "+
		"HasImages: %t \n   "+
		"Text: %s \n   "+
		"Images: %v",
		data.Author, data.Avatar, data.Timestamp, data.IsAReply, data.HasText, data.HasImages, data.Text, data.Images)
}

// String converts Quote properties to a string
func (quote Quote) String() string {
	return fmt.Sprintf("\n --- Quote ---\n %s", quote.Data.String())
}
