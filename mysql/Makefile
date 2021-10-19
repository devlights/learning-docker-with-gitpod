IMAGE_NAME=try-docker/mysql
IMAGE_TAG=latest

default: run

run: build
	docker container run \
		--detach \
		--interactive \
		--tty \
		--rm \
		--net net1 \
		-e MYSQL_ROOT_PASSWORD=rootroot \
		-e MYSQL_DATABASE=db1 \
		-e MYSQL_USER=user1 \
		-e MYSQL_PASSWORD=useruser \
		$(IMAGE_NAME):$(IMAGE_TAG) \
		--character-set-server=utf8mb4 \
		--collation-server=utf8mb4_unicode_ci \
		--default-authentication-plugin=mysql_native_password

build:
	docker network create net1
	docker container build -f Dockerfile -t $(IMAGE_NAME):$(IMAGE_TAG) .
