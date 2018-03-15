package gokolonial

import (
	"net/url"
	"net/http"
	"os"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	Token     string

	httpClient *http.Client

	cookie string
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewClient() (err error, client Client) {
	url, err := url.Parse("https://kolonial.no")
	if err != nil {
		return err, client
	}
	client = Client{BaseURL: url,
		UserAgent: os.Getenv("kolonialUseragent"),
		Token: os.Getenv("kolonialToken"),
		httpClient: http.DefaultClient}
	return err, client
}


