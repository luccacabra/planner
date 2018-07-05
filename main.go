package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	configFile = kingpin.Flag("config.file", "Planner configuration file.").Default("planner.yaml").String()
)

func main() {
	// Load config
	kingpin.Parse()
	baseName := filepath.Base(*configFile)
	viper.SetConfigName(strings.TrimSuffix(baseName, filepath.Ext(baseName)))
	viper.AddConfigPath(filepath.Dir(*configFile))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
