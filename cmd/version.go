// Copyright © 2016 Harald Ringvold <harald.ringvold@gmail.com>
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
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/kardianos/osext"
)

const VersionNumber = 0.2
const VersionSuffix = "-DEV" // blank this when doing a releas

var (
	CommitHash string
	BuildDate  string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the verson number",
	Long: `Shows the version. Pretty simple.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if BuildDate == "" {
			setBuildDate() // set the build date from executable's mdate
		} else {
			formatBuildDate() // format the compile time
		}
		if CommitHash == "" {
			fmt.Printf("Enonicstatus v%s BuildDate: %s\n", EnonicstatusVersion(), BuildDate)
		} else {
			fmt.Printf("Enonicstatus v%s-%s BuildDate: %s\n", EnonicstatusVersion(), strings.ToUpper(CommitHash), BuildDate)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

// setBuildDate checks the ModTime of the Hugo executable and returns it as a
// formatted string.  This assumes that the executable name is Hugo, if it does
// not exist, an empty string will be returned.  This is only called if the
// BuildDate wasn't set during compile time.
//
// osext is used for cross-platform.
// Code from Hugo https://github.com/spf13/hugo/blob/master/commands/version.go
func setBuildDate() {
	fname, _ := osext.Executable()
	dir, err := filepath.Abs(filepath.Dir(fname))
	if err != nil {
		fmt.Println(err)
		return
	}
	fi, err := os.Lstat(filepath.Join(dir, filepath.Base(fname)))
	if err != nil {
		fmt.Println(err)
		return
	}
	t := fi.ModTime()
	BuildDate = t.Format(time.RFC3339)
}

// formatBuildDate formats the BuildDate according to the value in
// .Params.DateFormat, if it's set.
func formatBuildDate() {
	t, _ := time.Parse("2006-01-02T15:04:05-0700", BuildDate)
	BuildDate = t.Format(time.RFC3339)
}

func EnonicstatusVersion() string {
	return enonicstatusVersion(VersionNumber, VersionSuffix)
}

func enonicstatusVersion(version float32, suffix string) string {
	return fmt.Sprintf("%.2g%s", version, suffix)
}

func enonicstatusVersionNoSuffix(version float32) string {
	return fmt.Sprintf("%.2g", version)
}
