package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

func getDockerHubToken(username, password, url string) *DockerHubClient {
	body := []byte(fmt.Sprintf("username=%s&password=%s", username, password))
	log.Println("Reaching URL : " + url + "/v2/users/login/")
	response, err := http.Post(url+"/v2/users/login/", "application/x-www-form-urlencoded", bytes.NewBuffer(body))

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

	token.Username = username
	token.Url = url
	log.Println("OK!")
	return &token
}

func getTagList(imageName string, client *DockerHubClient) DockerHubTagListResponse {
	httpClient := http.Client{}
	url := client.Url + fmt.Sprintf("/v2/repositories/%s/tags/", imageName)
	log.Printf("Getting tags list for %s. %s", imageName, url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "JWT "+client.Token)
	response, err := httpClient.Do(req)

	if err != nil || response == nil {
		log.Fatalln(errors.Wrap(err, fmt.Sprintf("Failed getting %s tags list", imageName)))
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var toReturn DockerHubTagListResponse
	err = decoder.Decode(&toReturn)

	if err != nil {
		return DockerHubTagListResponse{}
	}
	log.Println("OK!")
	return toReturn
}

func getLastTag(tags []DockerHubTag) string {
	for _, tag := range tags {
		if testExcludedTag(tag.Name) {
			return tag.Name
		}
	}
	return "latest"
}

func testExcludedTag(tag string) bool  {
	for _, excludedTag := range excludedTags {
		if excludedTag == tag {
			return false
		}
	}

	return true
}
