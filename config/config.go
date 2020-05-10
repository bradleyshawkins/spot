package config

import (
	"fmt"
	"os"

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
)

type SpotifyClientInfo struct {
	ClientID string
	Secret   string
}

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
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}

func GetSpotifyClientInfo() SpotifyClientInfo {
	return SpotifyClientInfo{
		ClientID: viper.GetString(spotifyClient),
		Secret:   viper.GetString(spotifySecret),
	}
}

func SetOAuthToken(token *oauth2.Token) error {

	viper.Set(spotifyAccessToken, token.AccessToken)
	viper.Set(spotifyRefreshToken, token.RefreshToken)
	viper.Set(spotifyAccessTokenExpiration, token.Expiry)

	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
