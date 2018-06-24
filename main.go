package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/spf13/viper"
)

var (
	configFile = kingpin.Flag("config.file", "ECR exporter configuration file.").Default("prometheus-to-datadog.yaml").String()
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