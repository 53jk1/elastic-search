VERSION = 1.0.0
COMMIT = $(shell git rev-parse --short HEAD)
BUILD_TIME = $(shell date +%FT%T%z)
LDFLAGS = -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildTime=$(BUILD_TIME)"
DATE = $(shell date +%Y%m%d)
IMAGE = $(shell basename $(shell pwd))
DOCKER_ELASTIC_SEARCH_IMAGE = docker.elastic.co/elasticsearch/elasticsearch:7.6.2

build:
	go build -o bin/$(NAME) -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

run:
	go run main.go

run_elastic:
	docker run -d --name elastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" $(DOCKER_ELASTIC_SEARCH_IMAGE)

stop_elastic:
	docker stop elastic
	docker rm elastic