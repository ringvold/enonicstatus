package cmd

import (
	"testing"

	"github.com/spf13/viper"
)

func setup() {
	if testing.Verbose() {
		debugEnabled = true
	}
}

func initViper() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".enonicstatus") // name of config file (without extension)
	viper.AddConfigPath("..") // adding current directory as first search path
  viper.AddConfigPath("$HOME")  //  home directory
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		Debug("Using config file:", viper.ConfigFileUsed())
	}
}

func TestGetPath(t *testing.T) {
	setup()

	got := GetPath("")
	want := "/status"
	if got != want {
		t.Errorf("GetPath() == %q, want %q", got, want)
	}
	initViper()

	got = GetPath("")
	want = "/enoniccms4.7.json"
	if got != want {
		t.Errorf("GetPath() == %q, want %q", got, want)
	}

	jsonPath = "changed"
	got = GetPath("")
	if got != "changed" {
		t.Error("Not using jsonPath flag")
	}
	jsonPath = "/status"


	// Get host for spesified enviroment
	env := "utv"
	got = GetPath(env)
	want = "/utvpath"
	if got != want {
		t.Errorf("GetPath(%q) == %q, want %q", env, got, want)
	}

	env = "doesnotexist"
	got = GetPath(env)
	want = jsonPath
	if got != want {
		t.Errorf("GetPath(%q) == %q, want %q", env, got, want)
	}

	viper.Reset()
}

func TestGetHosts(t *testing.T) {
	setup()

	got := GetHosts("")
	want := ""
	if got != want {
		t.Errorf("GetHosts() == %q, want %q", got, want)
	}

	initViper()
	got = GetHosts("")
	want = "localhost:2015,localhost:2015"
	if got != want {
		t.Errorf("GetHosts() == %q, want %q", got, want)
	}

	hosts = "other:8080,hosts:8080"
	got = GetHosts("")
	if got != hosts {
		t.Errorf("GetHosts() == %q, want %q", got, hosts)
	}
	hosts=""

	// Get host for spesified enviroment
	env := "utv"
	got = GetHosts(env)
	want = "utv.localhost:2016"
	if got != want {
		t.Errorf("GetHosts(%q) == %q, want %q", env, got, want)
	}

	env = "doesnotexist"
	got = GetHosts(env)
	want = ""
	if got != want {
		t.Errorf("GetHosts(%q) == %q, want %q", env, got, want)
	}

	viper.Reset()
}
