# Vars
VERSION=1.0
IMG_NAME=sproxy
REPO=lordwire
FULLNAME=${REPO}/${IMG_NAME}:${VERSION}
pwd=$(shell pwd)


## ----------------------------------------------------
## Sproxy: a simple reverse proxy implementation in Go. 
## 
## Requirements:
## - make
## - golang 1.17
## - docker
## - docker-compose
## - helm (tested with version 3.8.0)
## - a working kubernetes cluster
## ----------------------------------------------------
## 


.DEFAULT_GOAL := help

## Targets: 
## 
image: ## build an image using Docker (Go app is also built inside the image)
	@docker build . -t ${FULLNAME}

push: ## push the image to Docker Hub (Warning: requires Nick's credentials for this particular image)
	@docker push ${FULLNAME}

build: ## build a go binary locally. 
	@go build 

dockerun: ## run the container locally. Address is hardcoded to localhost:8080
	@docker run -p 8080:8080 --mount type=bind,source="$(pwd)"/config.yml,target=/app/config.yml ${FULLNAME}

helmon: ## run a helm install
	@helm install sproxy sproxy/

helmoff: ## run a helm delete
	@helm delete sproxy

test: ## Spin-up a test env with httpbin and sproxy
	@docker-compose up -d
	@curl --parallel  --parallel-max 300 --config testhosts/curl1000.txt

notest: ## bringdown the test environment
	@docker-compose down

help: ## print this help file
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)
