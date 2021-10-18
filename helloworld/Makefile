IMAGE_NAME=try-docker/helloworld
IMAGE_TAG=latest

default: run

run: build
	docker run --rm $(IMAGE_NAME):$(IMAGE_TAG)

build:
	docker build -f Dockerfile -t $(IMAGE_NAME):$(IMAGE_TAG) .

