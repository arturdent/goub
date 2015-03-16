package goub

import "golang.org/x/oauth2"

func GetOauthCodeLink(config *oauth2.Config, state string, accessType oauth2.AuthCodeOption) string {
	return config.AuthCodeURL(state, accessType)
}

func GetTokenFromCode(config *oauth2.Config, code string) (*oauth2.Token, error) {
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}
