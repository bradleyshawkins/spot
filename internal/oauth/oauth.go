package oauth

import (
	"fmt"

	"golang.org/x/oauth2"

	"github.com/zmb3/spotify"
)

var scopesRequested = []string{}

type OAuth struct {
	authenticator spotify.Authenticator
}

type oauthTokenResponse struct {
	token *oauth2.Token
	error error
}

func NewOAuth(clientID, clientSecret string) OAuth {
	a := spotify.NewAuthenticator(RedirectURL, scopesRequested...)
	a.SetAuthInfo(clientID, clientSecret)
	return OAuth{
		authenticator: a,
	}
}

func (o OAuth) Authorize() (*oauth2.Token, error) {
	tokenChan := make(chan oauthTokenResponse)
	go o.Callback(tokenChan)

	url := o.authenticator.AuthURL("")
	fmt.Println("Visit this URL to authorize", url)

	// Start http server for callback
	resp := <-tokenChan
	if resp.error != nil {
		return nil, resp.error
	}

	close(tokenChan)

	return resp.token, nil
}
