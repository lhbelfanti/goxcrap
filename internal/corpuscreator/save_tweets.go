package corpuscreator

import (
	"context"
	"fmt"

	"goxcrap/internal/http"
	"goxcrap/internal/log"
)

// SaveTweets calls the endpoint in charge of inserting into database the obtained tweets. It sends as body the slice of
// tweets retrieved with the scrapper
type SaveTweets func(ctx context.Context, body SaveTweetsBody) error

// MakeSaveTweets create a new SaveTweets
func MakeSaveTweets(httpClient http.Client, domain string) SaveTweets {
	url := domain + "/tweets/v1"

	return func(ctx context.Context, body SaveTweetsBody) error {
		resp, err := httpClient.NewRequest(ctx, "POST", url, body)
		if err != nil {
			log.Error(ctx, err.Error())
			return FailedToExecuteRequest
		}

		log.Info(ctx, fmt.Sprintf("Save tweets endpoint called -> Status: %s | Response: %s", resp.Status, resp.Body))

		return nil
	}
}
