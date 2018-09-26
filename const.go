package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)

const DockerHubApiUrl = "https://hub.docker.com"
var excludedTags = []string{"latest", "develop", "edge", "snapshot"}

func getConfig() *AppConfig {
	var opts AppConfig
	_, err := flags.Parse(&opts)

	// if some argument not provided
	if err != nil {
		os.Exit(1)
	}
	
	// if asking --help. https://github.com/jessevdk/go-flags/pull/263
	if flags.WroteHelp(err) {
		os.Exit(0)
	}
	
	return &opts
}
