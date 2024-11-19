package corpuscreator

type (
	// SaveTweetsBody represents the body for the Save Tweets endpoint call
	SaveTweetsBody []TweetDTO

	// TweetDTO represents the tweet to be saved
	TweetDTO struct {
		Hash             *string   `json:"hash,omitempty"`
		Author           string    `json:"author"`
		Avatar           *string   `json:"avatar,omitempty"`
		PostedAt         string    `json:"posted_at"`
		IsAReply         bool      `json:"is_a_reply"`
		TextContent      *string   `json:"text_content,omitempty"`
		Images           []string  `json:"images,omitempty"`
		Quote            *QuoteDTO `json:"quote,omitempty"`
		SearchCriteriaID *int      `json:"search_criteria_id,omitempty"`
	}

	// QuoteDTO represents the quote of the tweet to be saved
	QuoteDTO struct {
		Author      string   `json:"author"`
		Avatar      *string  `json:"avatar,omitempty"`
		PostedAt    string   `json:"posted_at"`
		IsAReply    bool     `json:"is_a_reply"`
		TextContent *string  `json:"text_content,omitempty"`
		Images      []string `json:"images,omitempty"`
	}

	// Execution represents the response for the Get Search Criteria Execution endpoint call
	Execution struct {
		ID               int    `json:"id"`
		Status           string `json:"status"`
		SearchCriteriaID int    `json:"search_criteria_id"`
	}

	// UpdateSearchCriteriaExecutionBody represents the body for the Update Search Criteria Execution endpoint call
	UpdateSearchCriteriaExecutionBody struct {
		Status string `json:"status"`
	}

	// InsertSearchCriteriaExecutionDayBody represents the body for the Insert Search Criteria Execution Day endpoint call
	InsertSearchCriteriaExecutionDayBody struct {
		ExecutionDate             string  `json:"execution_date"`
		TweetsQuantity            int     `json:"tweets_quantity"`
		ErrorReason               *string `json:"error_reason"`
		SearchCriteriaExecutionID int     `json:"search_criteria_execution_id"`
	}
)

const (
	InProgressStatus string = "IN PROGRESS"
	DoneStatus       string = "DONE"
)

// NewUpdateExecutionBody creates a new UpdateSearchCriteriaExecutionBody with the given status
func NewUpdateExecutionBody(status string) UpdateSearchCriteriaExecutionBody {
	return UpdateSearchCriteriaExecutionBody{
		Status: status,
	}
}

// NewInsertExecutionDayBody creates a new InsertSearchCriteriaExecutionDayBody with the given parameters
func NewInsertExecutionDayBody(executionDate string, tweetsQuantity int, errorReason *string, executionID int) InsertSearchCriteriaExecutionDayBody {
	return InsertSearchCriteriaExecutionDayBody{
		ExecutionDate:             executionDate,
		TweetsQuantity:            tweetsQuantity,
		ErrorReason:               errorReason,
		SearchCriteriaExecutionID: executionID,
	}
}
