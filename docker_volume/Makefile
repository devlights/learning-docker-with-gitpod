make-volume:
	docker volume create vol-1

rm-volume:
	docker volume rm vol-1

sh1:
	docker container run \
		--interactive \
		--tty \
		--rm \
		--volume vol-1:/app \
		--volume ${PWD}:/app/data \
		--workdir /app \
		debian:bullseye bash

sh2:
	docker container run \
		--interactive \
		--tty \
		--rm \
		--volume vol-1:/app \
		--workdir /app \
		debian:bullseye bash
