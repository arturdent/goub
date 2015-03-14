package goub

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	goubVersion         = "0.1"
	defaultBaseURL      = "http://coub.com"
	userAgent           = "goub/" + goubVersion
	authEndpoint        = "/oauth/authorize"
	accessTokenParametr = "access_token"
)

type Client struct {
	Client      *http.Client
	BaseURL     *url.URL
	UserAgent   string
	AccessToken string
}

func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	client := &Client{
		Client:      httpClient,
		BaseURL:     baseURL,
		UserAgent:   userAgent,
		AccessToken: token,
	}

	return client
}

func (client *Client) NewRequest(method, urlString string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	url := client.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}

	// req.Header.Add("Accept", "mediaType")
	if client.UserAgent != "" {
		req.Header.Add("User-Agent", client.UserAgent)
	}
	return req, nil
}
