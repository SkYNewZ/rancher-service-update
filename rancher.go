package main

import (
	"fmt"
	"github.com/pkg/errors"
	rancher "github.com/rancher/go-rancher/v2"
	"log"
	"regexp"
	"strings"
	"time"
)

func createClient(rancherURL, accessKey, secretKey string) (*rancher.RancherClient) {
	log.Println("Authenticating through Rancher Server...")
	client, err := rancher.NewRancherClient(&rancher.ClientOpts{
		Url:       rancherURL,
		AccessKey: accessKey,
		SecretKey: secretKey,
		Timeout:   time.Second * 20,
	})

	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed to create a client for rancher"))
	}

	log.Println("OK!")
	return client
}

func getRancherServicesList(client rancher.RancherClient) []Stack {
	log.Println("Getting services list...")
	m := map[string]interface{}{"system": false}
	options := &rancher.ListOpts{Filters: m}

	// Get all stacks
	stacks, _ := client.Stack.List(options)
	var finalList []Stack
	var countServices int
	for _, stack := range stacks.Data {
		// all services id in this stack
		var serviceList []Service
		for _, serviceId := range stack.ServiceIds {
			// get single service info
			service, _ := client.Service.ById(serviceId)
			image, _ := service.Data["fields"].(map[string]interface{})
			image = image["launchConfig"].(map[string]interface{})
			formattedImageName := formattedImage(image["imageUuid"].(string))
			serviceList = append(serviceList,
				Service{
					ImageName:  fmt.Sprintf("%s/%s", formattedImageName.Username, formattedImageName.ImageName),
					CurrentTag: formattedImageName.CurrentTag,
					Name:       service.Name,
				})
			countServices++
		}
		finalList = append(finalList, Stack{Name: stack.Name, Services: serviceList})
	}
	log.Println("OK!")
	log.Printf("Found %d stacks and %d services", len(finalList), countServices)
	return finalList
}

func formattedImage(originalImageName string) ImageRegexMatch {
	// remove `docker:`
	originalImageName = strings.Replace(originalImageName, "docker:", "", 1)

	// execute regex for getting correct part of image
	re := regexp.MustCompile(`(?m)[a-z0-9]+(?:[._-][a-z0-9]+)*`)
	result := re.FindAllString(originalImageName, -1)

	if !strings.Contains(originalImageName, ":") { // if image does not contain tag
		result = append(result, "latest")
	}

	switch len(result) {
	case 5: // custom registry with custom port. e.g cloud.canister.io:5000/skynewz/parrot-front:1.0
		return ImageRegexMatch{Username: fmt.Sprintf("%s:%s/%s", result[0], result[1], result[2]), ImageName: result[3], CurrentTag: result[4]}
	case 4: // custom registry. e.g registry.gitlab.com/skynewz/website:1.5
		return ImageRegexMatch{Username: result[0] + "/" + result[1], ImageName: result[2], CurrentTag: result[3]}
	case 3: // docker hub image. e.g skynewz/website:1.5
		return ImageRegexMatch{Username: result[0], ImageName: result[1], CurrentTag: result[2]}
	case 2: // official docker image postgres:10 -> library/postgres:10
		return ImageRegexMatch{Username: "library", ImageName: result[0], CurrentTag: result[1]}
	default:
		return ImageRegexMatch{Username: "", ImageName: "", CurrentTag: ""}
	}
}
