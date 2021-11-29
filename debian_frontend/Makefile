run:
	docker container run \
		--init \
		--interactive \
		--tty \
		--rm \
		mcr.microsoft.com/vscode/devcontainers/go:1.17 \
		sh -c "apt-get update -q; DEBIAN_FRONTEND=noninteractive apt-get install -yq vim; bash"
