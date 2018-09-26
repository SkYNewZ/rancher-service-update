[![Build Status](https://travis-ci.org/SkYNewZ/rancher-service-update.svg?branch=master)](https://travis-ci.org/SkYNewZ/rancher-service-update)

# Rancher service updates

This script written in Go will get all your rancher server services and compare with the latest tag on [DockerHub](https://hub.docker.com)

## Getting started

1. Download [latest binary release](https://github.com/SkYNewZ/rancher-service-update/releases/latest)
2. (Optional) Add into your `$PATH`
3. You need to have some required variables :
  - DockerHub username
  - DockerHub password
  - Rancher API access & secret keys (check http://[SERVER_URL]:[PORT]/env/[ENV]/api/keys)

```bash
$ rancher-service-update -u=username -p=password -s=http://[SERVER_URL]:[PORT] --access-key=access_key --secret-key=secret_key # Or via environment 
```

This will print by default a **table** like :

| STACK NAME | SERVICE NAME | IMAGE | ACTUAL TAG | LATEST TAG |
| :--------: | :----------: | :---: | :--------: | :--------: |
| Stack 1 | api | php | 7.2-stretch | 7.7-stretch |
| | front | registry/front | 1.1.4 | 1.1.5 |
| | database | library/postgres | 10.3-alpine | 10 |
| Other Stack | hello-world | tutum/hello-world | latest | latest |
| Another Stack | hello-world | ... | ... | ... |

> If you want to print result as JSON, add the option `-o=json`. The default is `-o=table`

## Run as Go binary
```bash
$ go get github.com/SkYNewZ/rancher-service-update
# Ensure that $GOPATH is set, or if $HOME/go/bin is into your $PATH
$ go-rancher-service-update -u=username -p=password -s=http://[SERVER_URL]:[PORT] --access-key=access_key --secret-key=secret_key # Or via environment variables
```

## Run as docker container

```bash
$ docker run -it --rm \
    -e DOCKER_USERNAME=username \
    -e DOCKER_PASSWORD=password \
    -e RANCHER_SERVER_URL=http://[SERVER_URL]:[PORT] \
    -e RANCHER_ACCESS_KEY=access_key \
    -e RANCHER_SECRET_KEY=secret_key \
    skynewz/rancher-service-update
```

## Usage
```
Usage:
  rancher-services-check-updates [OPTIONS]

Application Options:
  -u, --docker-username=    DockerHub username [$DOCKER_USERNAME]
  -p, --docker-password=    DockerHub password [$DOCKER_PASSWORD]
  -s, --server=             Rancher server URL [$RANCHER_SERVER_URL]
      --access-key=         Rancher API access key [$RANCHER_ACCESS_KEY]
      --secret-key=         Rancher API secret key [$RANCHER_SECRET_KEY]
  -o, --output=[json|table] Output type (default: table)

Help Options:
  -h, --help                Show this help message
```
