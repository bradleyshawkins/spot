package oauth

import (
	"fmt"

	"golang.org/x/oauth2"

	"github.com/zmb3/spotify"
)

var scopesRequested = []string{}

type OAuth struct {
	conf *oauth2.Config
}

type oauthTokenResponse struct {
	token *oauth2.Token
	error error
}

func NewOAuth(conf *oauth2.Config) OAuth {
	a := spotify.NewAuthenticator(conf.RedirectURL, scopesRequested...)
	a.SetAuthInfo(conf.ClientID, conf.ClientSecret)
	return OAuth{
		conf: conf,
	}
}

func (o OAuth) Authorize() (*oauth2.Token, error) {
	tokenChan := make(chan oauthTokenResponse)
	go o.Callback(tokenChan)

	url := o.conf.AuthCodeURL("", oauth2.AccessTypeOffline)
	fmt.Println("Visit this URL to authorize", url)

	// Start http server for callback
	resp := <-tokenChan
	if resp.error != nil {
		return nil, resp.error
	}

	close(tokenChan)

	return resp.token, nil
}
