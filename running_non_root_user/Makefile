ARG = go version

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
		--user 1000:1000 \
		--mount type=bind,src=${PWD},dst=/workspace \
		--workdir /workspace \
		--env GOCACHE=/workspace/.cache \
		mcr.microsoft.com/vscode/devcontainers/go:1.17 \
		${ARG}
