package config

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/zmb3/spotify"

	"github.com/pkg/errors"

	"golang.org/x/oauth2"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	spotifyClient                = "SPOTIFY_CLIENT"
	spotifySecret                = "SPOTIFY_SECRET"
	spotifyAccessToken           = "SPOTIFY_ACCESS_TOKEN"
	spotifyRefreshToken          = "SPOTIFY_REFRESH_TOKEN"
	spotifyAccessTokenExpiration = "SPOTIFY_ACCESS_TOKEN_EXPIRATION"
	spotifyTokenType             = "SPOTIFY_TOKEN_TYPE"

	SpotifyURL = "https://api.spotify.com"
)

var (
	spotifyRedirectURL = "http://localhost:8080/oauth/callback"
)

func InitConfig(cfgFile string) func() {
	return func() {
		if cfgFile != "" {
			// Use config file from the flag.
			viper.SetConfigFile(cfgFile)
		} else {
			// Find home directory.
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Search config in ~/.spot with name "config" (without extension).
			dir := home + "/.spot"
			viper.AddConfigPath(dir)
			viper.SetConfigName("config")
		}

		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error loading config file. Error:", err)
		}
	}
}

func GetClient(ctx context.Context) (*http.Client, error) {
	conf := GetOAuthConfig()
	token := GetOAuthToken()

	if token.Expiry.Before(time.Now()) {
		t, err := conf.TokenSource(ctx, token).Token()
		if err != nil {
			return nil, errors.Wrap(err, "unable to refresh token")
		}

		err = SetOAuthToken(t)
		if err != nil {
			return nil, errors.Wrap(err, "unable to set oauth token")
		}
	}

	return conf.Client(ctx, token), nil
}

func GetOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     viper.GetString(spotifyClient),
		ClientSecret: viper.GetString(spotifySecret),
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.spotify.com/authorize",
			TokenURL:  "https://accounts.spotify.com/api/token",
			AuthStyle: 0,
		},
		RedirectURL: spotifyRedirectURL,
		Scopes: []string{
			spotify.ScopeUserReadCurrentlyPlaying,
			spotify.ScopeUserFollowRead,
			spotify.ScopeUserFollowModify,
			spotify.ScopeUserModifyPlaybackState,
		},
	}
}

func SetOAuthToken(token *oauth2.Token) error {

	viper.Set(spotifyAccessToken, token.AccessToken)
	viper.Set(spotifyRefreshToken, token.RefreshToken)
	viper.Set(spotifyAccessTokenExpiration, token.Expiry)
	viper.Set(spotifyTokenType, token.TokenType)

	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func GetOAuthToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  viper.GetString(spotifyAccessToken),
		TokenType:    viper.GetString(spotifyTokenType),
		RefreshToken: viper.GetString(spotifyRefreshToken),
		Expiry:       viper.GetTime(spotifyAccessTokenExpiration),
	}
}
