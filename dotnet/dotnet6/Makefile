UID = $(shell id -u)
GID = ${UID}

build:
	docker image build \
		-t dotnet6 \
		-f Dockerfile \
		--build-arg USER_UID=${UID} \
		--build-arg USER_GID=${GID} \
		${PWD}

run: build
	docker container run \
		--interactive \
		--tty \
		--rm \
		--mount type=bind,src=${PWD},dst=/workspace \
		--workdir /workspace \
		dotnet6 \
		dotnet --info