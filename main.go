package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/olekukonko/tablewriter"
)

func main() {

	parser := flags.NewParser(&AppConfig, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}
	config := &AppConfig

	// Configure Rancher SDK
	rancherClient := createClient(config.RancherServer, config.RancherAccessKey, config.RancherSecretKey)
	// Get Docker Hub token
	dockerHubToken := getDockerHubToken(config.DockerUsername, config.DockerPassword)

	// List rancher services
	rancherServices := getRancherServicesList(*rancherClient)
	unwantedRegistries := []string{"registry.gitlab.com", "cloud.canister.io"}

	for _, stack := range rancherServices {
		for i := 0; i < len(stack.Services); i++ {
			service := &stack.Services[i]
			if byPassRegistry(service.ImageName, unwantedRegistries) {
				service.LatestTag = "Private registry ?"
			} else {
				service.LatestTag = getLastestTag(getTagList(service.ImageName, *dockerHubToken))
			}
		}
	}

	switch config.Output {
	case "json":
		printAsJSONgo(rancherServices)
	case "table":
		printTable(rancherServices)
	default:
		printTable(rancherServices)
	}
}

func byPassRegistry(a string, list []string) bool {
	for _, registry := range list {
		if strings.Contains(a, registry) {
			return true
		}
	}
	return false
}

func printAsJSONgo(stacks []Stack) {
	b, _ := json.Marshal(stacks)
	fmt.Println(string(b))
}

func printTable(stacks []Stack) {
	var data [][]string
	for _, stack := range stacks {
		displayStackName := true
		for _, service := range stack.Services {
			name := ""
			if displayStackName {
				name = stack.Name
			}
			data = append(data, []string{name, service.Name, service.ImageName, service.CurrentTag, service.LatestTag})
			displayStackName = false
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Stack name", "Service name", "Image", "Actual tag", "Latest tag"})
	table.SetBorder(true)
	table.AppendBulk(data)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}
