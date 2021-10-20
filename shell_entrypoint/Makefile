IMAGE_NAME=try-docker/shell
IMAGE_TAG=latest

default: run

run: build
	docker run -it --rm $(IMAGE_NAME):$(IMAGE_TAG)
	@true

build:
	docker build -f Dockerfile -t $(IMAGE_NAME):$(IMAGE_TAG) .

