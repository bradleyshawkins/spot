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

	"github.com/bradleyshawkins/spot/config"

	"github.com/spf13/cobra"
)

var expiration bool

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Get info about the oauth token in config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.GetOAuthToken()

		if expiration {
			fmt.Println("Expiration:\n\t", token.Expiry)
			return
		}
		fmt.Println("Access Token:\n\t", token.AccessToken)
		fmt.Println("Refresh Token:\n\t", token.RefreshToken)
		fmt.Println("Expiration:\n\t", token.Expiry)
		fmt.Println("Token Type:\n\t", token.TokenType)
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)

	tokenCmd.Flags().BoolVarP(&expiration, "expiration", "e", false, "Display only token expiration time")
}
