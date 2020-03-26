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
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var (
	remote string
	local  string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "sa",
		Short: "Service agent for relaying connections",
		Long: `Service agent faciliates making ssh and other tcp-based connections behind firewalls, by using a third-party external relay. 

Use sa serve at the service end, and sa connect at the client end.
The remote and local connections are convenient to specify via ENV
export SA_LOCAL=127.0.0.1:7799
export SA_REMOTE=wss://some.server.com:9988/some/route
sa serve

the remote destination should be fully explicit websocket URL e.g. wss://foo.server.com:9988/some/route
the local connection is by default 127.0.0.1:22; do not specify a protocol`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&local, "local", "l", "127.0.0.1:22", "local address (default is 127.0.0.1:22")
	rootCmd.PersistentFlags().StringVarP(&remote, "remote", "r", "", "remote address (required, no default")
	rootCmd.MarkFlagRequired("remote")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetEnvPrefix("SA")
	viper.AutomaticEnv() // read in environment variables that match
}
