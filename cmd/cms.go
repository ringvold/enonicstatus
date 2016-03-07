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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/haraldringvold/enonicstatus/jsonstruct"
	"github.com/haraldringvold/enonicstatus/formatter"
)

type GetJsonResult struct {
	json  jsonstruct.Status
	error error
}

const hostsViperPath = "hosts"
const jsonPathViperPath = "jsonPath"

// cmsCmd represents the status command
var CmsCmd = &cobra.Command{
	Use:   "cms",
	Short: "Shows status Enonic CMS nodes",
	Long:  `Extracts and diplays index status, uptime and master status for earch node`,
	RunE: func(cmd *cobra.Command, args []string) error {
		env := ""
		if len(args) != 0 {
			env = args[0]
		}

		path := GetPath(env)
		hosts := GetHosts(env)

		selectedFormatter := GetFormatter(viper.GetString(formatFlag))


		Debug("Path: ", path)
		Debug("Hosts: ", hosts)

		c := make(chan GetJsonResult)

		hostsSlice := strings.Split(hosts, ",")
		if hostsIsEpmty(hostsSlice) {
			return errors.New("No hosts configured")
		}

		for _, host := range hostsSlice {
			rawUrl := fmt.Sprintf("http://%v", host+path)
			hostUrl, err := url.Parse(rawUrl)
			if err != nil {
				panic(err)
			}
			go func() { c <- getJson(*hostUrl) }()
		}

		for i := 0; i < len(hostsSlice); i++ {
			select {
			case result := <-c:
				if result.error != nil {
					fmt.Println("")
					fmt.Println(result.error.Error())
				} else {
					printStatus(result.json, selectedFormatter)
				}
			}
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(CmsCmd)
}

func printStatus(json jsonstruct.Status, selectedFormatter formatter.Formatter) {
	fmt.Println("")
	fmt.Println(selectedFormatter.HostName(json.Cluster.LocalNode.HostName))
	fmt.Println(selectedFormatter.IndexStatus(json.Index.Status))
	fmt.Println(selectedFormatter.Master(json.Cluster.LocalNode.Master))
	fmt.Println(selectedFormatter.NodesSeen(json.Cluster.LocalNode.NumberOfNodesSeen))
	fmt.Println(selectedFormatter.Uptime(json.Jvm.UpTime))
	fmt.Println(selectedFormatter.Version(json.Product.Version))
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

func getJson(url url.URL) GetJsonResult {
	res := new(GetJsonResult)

	resp, err := http.Get(url.String())
	if err != nil {
		// TODO: Add debug statements?
		res.error = errors.New(fmt.Sprintf("Error: Could not connect to host %q. Is it correct?", &url))
		return *res
	}
	defer resp.Body.Close()

	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err)
	}

	var statusJson jsonstruct.Status

	if err := json.Unmarshal(body, &statusJson); err != nil {
		// TODO: Add debug statements
		res.error = errors.New(fmt.Sprintf("Cannot unmarshal json from host %q. Is it returning correct JSON?", &url))
	}
	res.json = statusJson
	return *res
}
