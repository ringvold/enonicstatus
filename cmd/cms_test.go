package cmd

import (
	"github.com/spf13/viper"
)

func ExamplePrintIndexStatus() {
	printIndexStatus("GREEN")
	// Output: # Index: [0;32;49m GREEN [0m
}

func ExamplePrintIndexStatusNoFormat() {
	viper.Set(noFormatingFlag, true)
	printIndexStatus("GREEN")
	// Output: # Index: GREEN
}
