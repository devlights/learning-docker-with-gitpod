UID = $(shell id -u)
GID = ${UID}
ARG = go build -o sample-app

default:

run-root:
	docker container run \
		--interactive \
		--tty \
		--rm \
		--mount type=bind,src=${PWD},dst=/workspace \
		--workdir /workspace \
		mcr.microsoft.com/vscode/devcontainers/go:1.17 \
		${ARG}

run-nonroot:
	docker container run \
		--interactive \
		--tty \
		--rm \
		--user ${UID}:${GID} \
		--mount type=bind,src=${PWD},dst=/workspace \
		--workdir /workspace \
		--env GOCACHE=/workspace/.cache \
		mcr.microsoft.com/vscode/devcontainers/go:1.17 \
		${ARG}

clean:
	rm -f sample-app