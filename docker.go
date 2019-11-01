package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const dockerHubAPIURL = "https://hub.docker.com"

var excludedTags = []string{"latest", "develop", "edge", "snapshot"}

func getDockerHubToken(username string, password string) *string {
	body := []byte(fmt.Sprintf("username=%s&password=%s", username, password))
	log.Println("Reaching URL : " + dockerHubAPIURL + "/v2/users/login/")
	response, err := http.Post(dockerHubAPIURL+"/v2/users/login/", "application/x-www-form-urlencoded", bytes.NewBuffer(body))

	if err != nil || response == nil {
		log.Fatalln(errors.Wrap(err, "Failed execute request on DockerHub API"))
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var token DockerHubClient
	err = decoder.Decode(&token)

	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed decode json response body on DockerHub API"))
	}
	log.Println("OK!")
	return &token.Token
}

func getTagList(imageName string, dockerHubToken string) *[]DockerHubTag {
	httpClient := http.Client{}
	url := dockerHubAPIURL + fmt.Sprintf("/v2/repositories/%s/tags/", imageName)
	log.Printf("Getting tags list for %s. %s", imageName, url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "JWT "+dockerHubToken)
	response, err := httpClient.Do(req)

	if err != nil || response == nil {
		log.Fatalln(errors.Wrap(err, fmt.Sprintf("Failed getting %s tags list", imageName)))
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var toReturn DockerHubTagListResponse
	err = decoder.Decode(&toReturn)

	if err != nil {
		return &[]DockerHubTag{}
	}
	log.Println("OK!")
	return &toReturn.Results
}

func getLastestTag(tags *[]DockerHubTag) string {
	for _, tag := range *tags {
		if !isExcludedImageTag(tag.Name) {
			return tag.Name
		}
	}
	return "latest"
}

func isExcludedImageTag(tag string) bool {
	for _, excludedTag := range excludedTags {
		return strings.Contains(tag, excludedTag)
	}
	return false
}
