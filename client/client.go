package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	HTTPClient *http.Client
	APIKey     string
	Req        *http.Request
	Headers    *map[string]string
}

func NewClient(apiKey *string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: time.Second * 10},
		APIKey:     *apiKey,
	}
}

func (c *Client) NewRequest(path, method, restURL string, body io.Reader) (*http.Request, error) {
	c.Req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", restURL, path), body)
	if err != nil {
		return nil, err
	}
	return c.Req, err
}

func (c *Client) DoRequest() ([]byte, error) {
	for key, value := range *c.Headers {
		c.Req.Header.Set(key, value)
	}
	res, err := c.HTTPClient.Do(c.Req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		return body, err
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}
