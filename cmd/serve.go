/*
Copyright © 2020 Tim Drysdale <timothy.d.drysdale@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/timdrysdale/grwc"
	"github.com/timdrysdale/tc"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Make a service available via a websocket client",
	Long:  `Connects to a local service port, and relays messages to a websocket relay server.`,
	Run: func(cmd *cobra.Command, args []string) {
		lcl := viper.GetString("local")
		rem := viper.GetString("remote")
		fmt.Printf("serve called with:\nlocal %s\nremote %s\n", lcl, rem)

		closed := make(chan struct{})

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		go func() {
			for _ = range c {
				close(closed)
			}
		}()

		// start the local connection

		lclConfig := &tc.Config{
			MaxFrameBytes: 1024000,
			Destination:   lcl,
		}

		tc := tc.New(lclConfig)
		go tc.Run(closed)

		// start the remote connection
		remConfig := grwc.Config{
			Destination:         rem,
			ExclusiveConnection: true, //force msgs to []byte on Receive channel
		}
		wc, err := grwc.New(&remConfig)

		if err != nil {
			log.Fatalf("Problem creating websocket client")
		}

		go wc.Run(closed)

		//connect messages from each port
		go func() {
		LOOP:
			for {
				select {
				case msg, ok := <-tc.Receive:
					if ok {
						wc.Send <- msg
					}
				case <-closed:
					break LOOP
				}
			}
		}()

		go func() {
		LOOP:
			for {
				select {
				case msg, ok := <-wc.Receive:
					if ok {
						tc.Send <- msg
					}
				case <-closed:
					break LOOP
				}
			}
		}()
		//wait for ctrl-c
		<-closed

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
