/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/bradleyshawkins/spot/internal/oauth"

	"github.com/bradleyshawkins/spot/config"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Spotify using OAuth2",
	Long: `Auth will return a URL to visit to authorize this app to
use the Spotify API. Visit the URL, click Agree and close the window.
The OAuth token will be stored in the specified config file. Default ~/.spot/config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {

		conf := config.GetOAuthConfig()

		auth := oauth.NewOAuth(conf)

		token, err := auth.Authorize()
		if err != nil {
			fmt.Println(err)
			return
		}

		err = config.SetOAuthToken(token)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
