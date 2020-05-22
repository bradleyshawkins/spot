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

	"github.com/bradleyshawkins/spot/config"
	"github.com/zmb3/spotify"

	"github.com/spf13/cobra"
)

// followCmd represents the follow command
var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Follows the artists of the currently playing song",
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

		var artistIDs []spotify.ID
		for _, artist := range fs.Item.Artists {
			artistIDs = append(artistIDs, artist.ID)
		}

		err = s.FollowArtist(artistIDs...)
		if err != nil {
			fmt.Println("Unable to follow artist/s. Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(followCmd)
}
