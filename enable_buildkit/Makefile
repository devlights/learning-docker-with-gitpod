use_env:
	DOCKER_BUILDKIT=1 docker image build -t use_env -f Dockerfile ${PWD}
	docker container run --rm use_env

use_buildx:
	docker image buildx build -t use_buildx -f Dockerfile ${PWD}
