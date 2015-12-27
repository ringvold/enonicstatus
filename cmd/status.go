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
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wsxiaoys/terminal/color"

	"github.com/haraldringvold/enonicstatus/jsonstruct"
)

var printLinePrefix string = "# "

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Shows status for configured nodes",
	Long:  `Extracts and diplays index status, uptime and master status for earch node`,
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("jsonPath")
		fmt.Println("Path: ", path)

		hosts := getHosts()
		for _, host := range hosts {
			printStatusForHost(host, path)
		}
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}

func printStatusForHost(host string, path string) {
	rawUrl := "http://"+host+"/"+path
	hostUrl, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}

	json := getJson(hostUrl)

	fmt.Println("")
	fmt.Println("Getting status for host:", host)
	fmt.Println("######")
	printName(json.Cluster.LocalNode.HostName)
	printIndexStatus(json.Index.Status)
	printMaster(json.Cluster.LocalNode.Master)
	printNodesSeen(json.Cluster.LocalNode.NumberOfNodesSeen)
	printUptime(json.Jvm.UpTime)
	fmt.Println("######")
}

func printName(name string) {
	fmt.Println(printLinePrefix+"Name:", name)
}
func printIndexStatus(status string) {
	formatting := ""
	if status == "GREEN" {
		formatting = "@g"
	}
	if status == "YELLOW" {
		formatting = "@y"
	}
	if status == "RED" {
		formatting = "@r"
	}
	color.Println(printLinePrefix+"Index:", formatting, status)
}

func printMaster(master string) {
	formatting := ""
	if master == "true" {
		formatting = "@g"
	}
	color.Println(printLinePrefix+"Master:", formatting, master)
}

func printUptime(uptime float64) {
	duration := fmt.Sprintf("%ims", uptime)
	formattedUptime, _ := time.ParseDuration(duration)
	formatting := "@b"
	color.Println(printLinePrefix+"Uptime:", formatting, formattedUptime)
}

func printNodesSeen(nodesSeen float64) {
	fmt.Println(printLinePrefix+"Nodes seen:", nodesSeen)
}

func getHosts() []string {
	return strings.Split(viper.GetString("hosts"), ",")
}

func getJson(url *url.URL) jsonstruct.Status {
	resp, err := http.Get(url.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err)
	}

	var statusJson jsonstruct.Status

	if err := json.Unmarshal(body, &statusJson); err != nil {
		panic(err)
	}

	return statusJson
}
