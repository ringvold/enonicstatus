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
	"strconv"
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wsxiaoys/terminal/color"

	"github.com/haraldringvold/enonicstatus/jsonstruct"
)

var hostsFlag string
var pathFlag string
var printLinePrefix string = "# "

const hostsViperPath = "hosts"
const jsonPathViperPath = "jsonPath"

// cmsCmd represents the status command
var CmsCmd = &cobra.Command{
	Use:   "cms",
	Short: "Shows status Enonic CMS nodes",
	Long:  `Extracts and diplays index status, uptime and master status for earch node`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path := viper.GetString("jsonPath")
		fmt.Println("Path: ", path)
		fmt.Println("Hosts: ", viper.GetString("hosts"))

		c := make(chan jsonstruct.Status)

		hosts := getHosts()
		if hostsIsEpmty(hosts) {
			return errors.New("No hosts configured")
		}

		for _, host := range hosts {
			rawUrl := "http://"+host+"/"+path
			hostUrl, err := url.Parse(rawUrl)
			if err != nil {
				panic(err)
			}
			go func() {c <- getJson(hostUrl) }()
		}

		for i := 0; i < len(hosts); i++ {
			select {
        case json := <-c:
					printStatus(json)
        }
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(CmsCmd)
}

func printStatus(json jsonstruct.Status) {
	fmt.Println("")
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
	uptimeString := strconv.FormatFloat(uptime, 'f', -1, 64)
	duration := fmt.Sprintf("%sms", uptimeString)
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

func hostsIsEpmty(hosts []string) bool {
	if len(hosts) < 0 {
		return true
	}
	if len(hosts) == 1 && hosts[0] == "" {
		return true
	}
	return false
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
