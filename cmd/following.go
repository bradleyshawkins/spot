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
	"strconv"

	"github.com/bradleyshawkins/spot/config"
	"github.com/zmb3/spotify"

	"github.com/spf13/cobra"
)

// followingCmd represents the following command
var followingCmd = &cobra.Command{
	Use:   "following",
	Short: "Checks to see if you follow the artists of the currently playing song",
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

		follows, err := s.CurrentUserFollows("artist", artistIDs...)
		if err != nil {
			fmt.Println("Error checking if user follows artists. Error:", err)
			return
		}

		// follows comes back in the order of the artistIDs sent in.
		for i, artist := range fs.Item.Artists {
			if follows[i] {
				check, _ := strconv.ParseInt(`2705`, 16, 32)
				fmt.Printf("%c  Following %s\n", rune(check), artist.Name)
			} else {
				x, _ := strconv.ParseInt(`274C`, 16, 32)
				fmt.Printf("%c  Not following %s\n", rune(x), artist.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(followingCmd)
}
