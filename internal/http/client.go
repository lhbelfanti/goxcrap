package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"goxcrap/internal/log"
)

type (
	// Client is an abstraction of the CustomClient methods
	Client interface {
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

// NewClient create a new CustomClient
func NewClient() *CustomClient {
	return &CustomClient{
		HTTPClient: &http.Client{Timeout: 3 * time.Second},
	}
}

func (c *CustomClient) NewRequest(ctx context.Context, method, url string, body interface{}) (Response, error) {
	ctx = log.With(ctx, log.Param("body", body), log.Param("url", url))

	var jsonData []byte
	var err error

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

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error(ctx, err.Error())
		return Response{}, FailedToCreateRequest
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Error(ctx, err.Error())
		return Response{}, FailedToExecuteRequest
	}
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
}
