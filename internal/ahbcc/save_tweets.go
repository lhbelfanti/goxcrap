package ahbcc

import (
	"fmt"
	"log/slog"

	"goxcrap/internal/http"
)

// SaveTweets calls the endpoint in charge of inserting into database the obtained tweets. It sends as body the slice of
// tweets retrieved with the scrapper
type SaveTweets func(body SaveTweetsBody) error

// MakeSaveTweets create a new SaveTweets
func MakeSaveTweets(httpClient http.Client, domain string) SaveTweets {
	url := domain + "/tweets/v1"

	return func(body SaveTweetsBody) error {
		resp, err := httpClient.NewRequest("POST", url, body)
		if err != nil {
			return FailedToExecuteRequest
		}

		slog.Info(fmt.Sprintf("Save tweets endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return nil
	}
}
