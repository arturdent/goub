package goub

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

const (
	goubVersion    = "0.1"
	defaultBaseURL = "http://coub.com"
	userAgent      = "goub/" + goubVersion
	authEndpoint   = "/oauth/authorize"
	authURL        = defaultBaseURL + "/oauth/authorize"
	tokenURL       = defaultBaseURL + "/oauth/token"
)

type Client struct {
	Client    *http.Client
	UserAgent string
	BaseURL   *url.URL
}

func MakeConfig(appId string, secret string, callbackURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     appId,
		ClientSecret: secret,
		RedirectURL:  callbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
	}
}

func NewClient(config *oauth2.Config, token *oauth2.Token) (*Client, error) {
	baseURL, _ := url.Parse(defaultBaseURL)

	client := &Client{
		Client:    config.Client(oauth2.NoContext, token),
		UserAgent: userAgent,
		BaseURL:   baseURL,
	}
	return client, nil
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

	return req, nil
}

func (cleint *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := cleint.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
