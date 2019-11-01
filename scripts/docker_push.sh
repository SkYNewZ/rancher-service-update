#!/bin/bash

# Login to registry
echo $GITHUB_TOKEN | docker login docker.pkg.github.com -u skynewz --password-stdin

# Prepare latest
docker tag rancher-service-update:latest docker.pkg.github.com/skynewz/rancher-service-update/rancher-service-update:latest
docker push docker.pkg.github.com/skynewz/rancher-service-update/rancher-service-update:latest

# Push tagged
docker tag rancher-service-update:latest docker.pkg.github.com/skynewz/rancher-service-update/rancher-service-update:$TRAVIS_TAG
docker push docker.pkg.github.com/skynewz/rancher-service-update/rancher-service-update:$TRAVIS_TAG
