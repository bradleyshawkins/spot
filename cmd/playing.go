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
	"context"

	"github.com/bradleyshawkins/spot/internal/spotify"

	"github.com/bradleyshawkins/spot/config"

	"github.com/spf13/cobra"
)

// playingCmd represents the playing command
var playingCmd = &cobra.Command{
	Use:   "playing",
	Short: "Gets currently playing song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		conf := config.GetOAuthConfig()
		token := config.GetOAuthToken()

		c := conf.Client(ctx, token)

		spotify.CurrentlyPlaying(c)

	},
}

func init() {
	rootCmd.AddCommand(playingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
