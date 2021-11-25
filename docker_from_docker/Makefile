default:

use-alpine-image:
	docker image build -t dfromd/alpine -f Dockerfile.alpine ${PWD}
	docker container run \
		--interactive \
		--tty \
		--rm \
		--mount type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock \
		dfromd/alpine \
		/bin/sh -c "docker image list"

use-debian-image:
	docker image build -t dfromd/debian -f Dockerfile.debian ${PWD}
	docker container run \
		--interactive \
		--tty \
		--rm \
		--mount type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock \
		dfromd/debian \
		/bin/sh -c "docker image list"	

use-docker-image:
	docker container run \
		--interactive \
		--tty \
		--rm \
		--mount type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock \
		docker:latest \
		docker image list
