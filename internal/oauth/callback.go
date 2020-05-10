package oauth

import (
	"context"
	"net/http"
	"sync"

	"github.com/opentracing/opentracing-go/log"
)

const (
	RedirectURL = "http://localhost:8080/oauth/callback"
)

func (o OAuth) Callback(codeChan chan<- oauthTokenResponse) {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// start http server
	svc := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/oauth/callback", o.callbackHandler(wg, codeChan))

	go func() {
		if err := svc.ListenAndServe(); err != http.ErrServerClosed {
			log.Error(err)
		}
	}()

	wg.Wait()

	// stop http server
	err := svc.Shutdown(context.TODO())
	if err != nil {

	}
}

func (o OAuth) callbackHandler(wg *sync.WaitGroup, tokenChan chan<- oauthTokenResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer wg.Done()

		token, err := o.authenticator.Token("", r)

		tokenChan <- oauthTokenResponse{
			token: token,
			error: err,
		}
	}
}
