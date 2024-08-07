package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type (
	// Client is an abstraction of the CustomClient methods
	Client interface {
		NewRequest(method, url string, body interface{}) (Response, error)
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

func (c *CustomClient) NewRequest(method, url string, body interface{}) (Response, error) {
	log.Debug().Msgf("Body: %#v, URL: %s", body, url)

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Error().Msg(err.Error())
		return Response{}, FailedToMarshalBody
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Msg(err.Error())
		return Response{}, FailedToCreateRequest
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Error().Msg(err.Error())
		return Response{}, FailedToExecuteRequest
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Error().Msg(err.Error())
		}
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msg(err.Error())
		return Response{}, FailedToReadResponse
	}

	return Response{
		Body:   string(respBody),
		Status: resp.Status,
	}, nil
}
