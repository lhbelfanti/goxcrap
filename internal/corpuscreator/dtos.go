package corpuscreator

type (
	// SaveTweetsBody represents the body for the Save Tweets endpoint call
	SaveTweetsBody []TweetDTO

	// TweetDTO represents the tweet to be saved
	TweetDTO struct {
		Hash             *string   `json:"hash,omitempty"`
		IsAReply         bool      `json:"is_a_reply"`
		TextContent      *string   `json:"text_content,omitempty"`
		Images           []string  `json:"images,omitempty"`
		Quote            *QuoteDTO `json:"quote,omitempty"`
		SearchCriteriaID *int      `json:"search_criteria_id,omitempty"`
	}

	// QuoteDTO represents the quote of the tweet to be saved
	QuoteDTO struct {
		IsAReply    bool     `json:"is_a_reply"`
		TextContent *string  `json:"text_content,omitempty"`
		Images      []string `json:"images,omitempty"`
	}
)
