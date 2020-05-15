package oauth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/oauth2"
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

		vals := r.URL.Query()

		if e := vals.Get("error"); e != "" {
			tokenChan <- oauthTokenResponse{
				token: nil,
				error: fmt.Errorf("error getting oauth token. Error: %v", e),
			}
		}

		code := vals.Get("code")
		if code == "" {
			tokenChan <- oauthTokenResponse{
				token: nil,
				error: errors.New("no code was returned"),
			}
		}

		token, err := o.conf.Exchange(r.Context(), code, oauth2.AccessTypeOffline)

		tokenChan <- oauthTokenResponse{
			token: token,
			error: err,
		}
	}
}
