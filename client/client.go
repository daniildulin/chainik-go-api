package client

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"os"
	"strings"
)

func (c *Client) TrackedCoin(symbol string) (TrackedToken, error) {
	list, err := c.TrackedCoins()
	if err != nil {
		return TrackedToken{}, err
	}

	for _, tc := range list {
		if strings.ToLower(tc.Ticker) == strings.ToLower(symbol) {
			return tc, nil
		}
	}

	return TrackedToken{}, errors.New("coin not found")
}

func (c *Client) TrackedCoins() ([]TrackedToken, error) {
	tcResponse := new(TrackedTokenResponse)
	e := new(ErrorResponse)
	resp, err := c.httpClient.R().
		SetResult(tcResponse).
		SetError(e).
		Get("/v2/tokens/tracked")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, e.Error
	}
	return tcResponse.Data, nil
}

type Client struct {
	httpClient *resty.Client
}

func NewClient() *Client {
	var client *resty.Client
	client = resty.New().SetBaseURL(os.Getenv("CHAINIK_API_HOST"))
	return &Client{httpClient: client}
}
