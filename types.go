package main

import "time"

// Service existant service in Rancher
type Service struct {
	ImageName  string
	CurrentTag string
	LatestTag  string
	Name       string
}

// Stack existant Stack in rancher
type Stack struct {
	Name     string
	Services []Service
}

// DockerHubTagListResponse List of tags
type DockerHubTagListResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []DockerHubTag `json:"results"`
}

// DockerHubTag Docker Hub tag
type DockerHubTag struct {
	Creator  int `json:"creator"`
	FullSize int `json:"full_size"`
	ID       int `json:"id"`
	ImageID  int `json:"image_id"`
	Images   []struct {
		Architecture string `json:"architecture"`
		Features     string `json:"features"`
		Os           string `json:"os"`
		OsFeatures   string `json:"os_features"`
		OsVersion    string `json:"os_version"`
		Size         int    `json:"size"`
		Variant      string `json:"variant"`
	} `json:"images"`
	LastUpdated time.Time `json:"last_updated"`
	LastUpdater int       `json:"last_updater"`
	Name        string    `json:"name"`
	Repository  int       `json:"repository"`
	V2          bool      `json:"v2"`
}

// DockerHubClient get Docker Hub API Token
type DockerHubClient struct {
	Token string `json:"token"`
}

// DockerImageWithTag create proper image
type DockerImageWithTag struct {
	Username   string
	ImageName  string
	CurrentTag string
}

// AppFlagsConfig main app config flags
type AppFlagsConfig struct {
	DockerUsername   string `short:"u" long:"docker-username" description:"DockerHub username" required:"true" env:"DOCKER_USERNAME"`
	DockerPassword   string `short:"p" long:"docker-password" description:"DockerHub password" required:"true" env:"DOCKER_PASSWORD"`
	RancherServer    string `short:"s" long:"server" description:"Rancher server URL" required:"true" env:"RANCHER_SERVER_URL"`
	RancherAccessKey string `long:"access-key" description:"Rancher API access key" required:"true" env:"RANCHER_ACCESS_KEY"`
	RancherSecretKey string `long:"secret-key" description:"Rancher API secret key" required:"true" env:"RANCHER_SECRET_KEY"`
	Output           string `short:"o" long:"output" description:"Output type" default:"table" choice:"json" choice:"table"`
}

// AppConfig main app config
var AppConfig AppFlagsConfig
