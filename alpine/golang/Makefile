UID = $(shell id -u)
GID = ${UID}

build:
	docker image build -t alpine/golang -f Dockerfile .

run: build
	docker container run \
		--interactive \
		--tty \
		--rm \
		--mount type=bind,src=${PWD},dst=/src \
		--workdir /src \
		--user ${UID}:${GID} \
		--env-file .env \
		alpine/golang \
		sh -c "go run ."