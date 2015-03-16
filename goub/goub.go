package goub

import (
	"net/http"

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

func NewClient(config *oauth2.Config, code string) (*Client, error) {
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Client:    config.Client(oauth2.NoContext, token),
		UserAgent: userAgent,
	}

	return client, nil
}

// func

// func (client *Client) NewRequest(method, urlString string, body interface{}) (*http.Request, error) {
// 	rel, err := url.Parse(urlString)
// 	if err != nil {
// 		return nil, err
// 	}

// 	url := client.BaseURL.ResolveReference(rel)

// 	var buf io.ReadWriter
// 	if body != nil {
// 		buf = new(bytes.Buffer)
// 		err := json.NewEncoder(buf).Encode(body)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	req, err := http.NewRequest(method, url.String(), buf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// req.Header.Add("Accept", "mediaType")
// 	if client.UserAgent != "" {
// 		req.Header.Add("User-Agent", client.UserAgent)
// 	}
// 	return req, nil
// }
