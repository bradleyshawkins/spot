package spotify

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bradleyshawkins/spot/config"
)

const (
	uri = "/v1/me/player/currently-playing"
)

func CurrentlyPlaying(c *http.Client) {
	req, err := http.NewRequest(http.MethodGet, config.SpotifyURL+uri, nil)
	if err != nil {

	}

	resp, err := c.Do(req)
	if err != nil {

	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	fmt.Println(string(b))
}
