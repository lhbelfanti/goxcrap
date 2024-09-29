package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"goxcrap/internal/log"
)

type (
	// Client is an abstraction of the CustomClient methods
	Client interface {
		// NewRequest creates a new HTTP Requests and executes it
		NewRequest(ctx context.Context, method, url string, body interface{}) (Response, error)
	}

	// CustomClient represent a custom http.CustomClient
	CustomClient struct {
		HTTPClient *http.Client
	}

	// Response represent the necessary data of the request response
	Response struct {
		Body   string
		Status string
	}
)

const maxRetries int = 3

// NewClient create a new CustomClient
func NewClient() *CustomClient {
	return &CustomClient{
		HTTPClient: &http.Client{Timeout: 3 * time.Second},
	}
}

// NewRequest creates a new HTTP Requests and executes it
func (c *CustomClient) NewRequest(ctx context.Context, method, url string, body interface{}) (Response, error) {
	var jsonData []byte
	var err error

	// Marshal body to JSON
	switch v := body.(type) {
	case []byte:
		jsonData = v
	default:
		jsonData, err = json.Marshal(body)
		if err != nil {
			log.Error(ctx, err.Error())
			return Response{}, FailedToMarshalBody
		}
	}

	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Create new HTTP request
		req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Error(ctx, err.Error())
			return Response{}, FailedToCreateRequest
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.HTTPClient.Do(req)
		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 30 {
			// Successful response: returning the data
			defer func(body io.ReadCloser) {
				err = body.Close()
				if err != nil {
					log.Error(ctx, err.Error())
				}
			}(resp.Body)

			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Error(ctx, err.Error())
				return Response{}, FailedToReadResponse
			}

			return Response{
				Body:   string(respBody),
				Status: resp.Status,
			}, nil
		} // else { // Unsuccessful response: retrying request }

		log.Error(ctx, fmt.Sprintf("Request to %s failed with error %v. \nRetrying... (Attempt %d/%d)", url, err, attempt, maxRetries))

		// Delay before retrying
		time.Sleep(100 * time.Millisecond)
	}

	log.Error(ctx, FailedToExecuteRequest.Error())
	return Response{}, FailedToExecuteRequest
}
