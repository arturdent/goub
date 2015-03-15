package goub

import "golang.org/x/oauth2"

func getOauthCodeLink(config oauth2.Config, state string, accessType oauth2.AuthCodeOption) {
	return config.AuthCodeURL(state, accessType)
}
