build:
	docker buildx build -t godev -f Dockerfile ${PWD}

run:
	docker container run \
		--init \
		--interactive \
		--tty \
		--rm \
		godev
