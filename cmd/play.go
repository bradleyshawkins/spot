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

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "plays a paused song",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		c, err := config.GetClient(ctx)
		if err != nil {
			fmt.Println("Error getting http client. Error:", err)
			return
		}

		s := spotify.NewClient(c)
		err = s.Play()
		if err != nil {
			fmt.Println("Error playing track. Error:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
