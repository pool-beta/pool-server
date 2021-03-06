package api

import (
	"io"
	"net/http"
)

type Context interface {
	Do(*http.Request) (*http.Response, error)

	Url() string
	Client() *http.Client
}

type context struct {
	url string
	client *http.Client
}

func NewContext(url string) (Context, error) {
	return &context{
		url: url,
		client: &http.Client{},
	}, nil
}

func (c *context) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

// Getters
func (c *context) Url() string {
	return c.url
}

func (c *context) Client() *http.Client {
	return c.client
}

// ----------------------------------- 	REQUEST ---------------------------------------------------

func NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}