package goub

import "golang.org/x/oauth2"

func GetOauthCodeLink(config *oauth2.Config, state string, accessType oauth2.AuthCodeOption) string {
	return config.AuthCodeURL(state, accessType)
}
