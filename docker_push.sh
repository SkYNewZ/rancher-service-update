#!/bin/bash

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push skynewz/rancher-service-update:latest
docker tag skynewz/rancher-service-update:latest skynewz/rancher-service-update:$TRAVIS_TAG
docker push skynewz/rancher-service-update:$TRAVIS_TAG