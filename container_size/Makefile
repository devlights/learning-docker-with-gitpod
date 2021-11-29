build:
	docker image build \
		-t myapp \
		-f Dockerfile \
		${PWD}

run: build
	docker container run \
		--init \
		--detach \
		--interactive \
		--tty \
		--rm \
		--name myapp1 \
		myapp
	docker container list --size
	docker container stop myapp1
	docker image rm myapp
	docker image prune -f
