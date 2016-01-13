// Copyright © 2015 Harald Ringvold <harald.ringvold@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	// "runtime/debug"
	// "errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debugEnabled bool

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "enonicstatus",
	Short: "Displays information about an Enonic CMS cluster",
	Long: `A commandline application that displays various information
about an Enonic cluster.

Enonicstatus gets the information from the status json that shows
information about the cluster and the current node.

Currently supports Enonic CMS with plans for Enonic XP.`,

// Uncomment the following line if your bare application
// has an action associated with it:
// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.enonicstatus.yaml)")

	RootCmd.PersistentFlags().String("hosts", "", "enonic nodes to check")
	RootCmd.PersistentFlags().String("json_path", "/status", "path on host to status json")
	RootCmd.PersistentFlags().BoolVar(&debugEnabled, "debug", false, "show more information on errors")
	viper.BindPFlag(hostsViperPath, RootCmd.PersistentFlags().Lookup("hosts"))
	viper.BindPFlag(jsonPathViperPath, RootCmd.PersistentFlags().Lookup("json_path"))
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".enonicstatus") // name of config file (without extension)
  viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AddConfigPath(".") // current directory
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
