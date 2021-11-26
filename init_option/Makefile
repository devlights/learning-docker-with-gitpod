build:
	docker image build -t init-option/app -f Dockerfile ${PWD}

no-init:
	docker container run \
		--detach \
		--interactive \
		--tty \
		--rm \
		--name container-no-init \
		init-option/app
	docker container exec -it container-no-init sh -c "curl http://localhost:8888/ ; ps aux"
	docker container stop container-no-init

with-init:
	docker container run \
		--init \
		--detach \
		--interactive \
		--tty \
		--rm \
		--name container-with-init \
		init-option/app
	docker container exec -it container-with-init sh -c "curl http://localhost:8888/ ; ps aux"
	docker container stop container-with-init