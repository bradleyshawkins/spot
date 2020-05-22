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
	"fmt"

	"github.com/zmb3/spotify"

	"github.com/bradleyshawkins/spot/config"

	"github.com/spf13/cobra"
)

// playingCmd represents the playing command
var playingCmd = &cobra.Command{
	Use:   "playing",
	Short: "Gets currently playing track",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		c, err := config.GetClient(ctx)
		if err != nil {
			fmt.Println("Error getting http client", err)
			return
		}

		s := spotify.NewClient(c)
		fs, err := s.PlayerCurrentlyPlaying()
		if err != nil {
			fmt.Println("Error getting currently playing song", err)
			return
		}

		fmt.Println("Track:\n\t", fs.Item.Name)

		if len(fs.Item.Artists) == 1 {
			fmt.Println("Artist:\n\t", fs.Item.Artists[0].Name)
		} else {
			fmt.Println("Artists:")
			for _, a := range fs.Item.Artists {
				fmt.Println("\t", a.Name)
			}
		}

		fmt.Println("Album:\n\t", fs.Item.Album.Name)
	},
}

func init() {
	rootCmd.AddCommand(playingCmd)
}
