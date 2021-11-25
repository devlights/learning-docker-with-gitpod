start:
	docker volume create vol-dind
	docker container run \
		--detach \
		--rm \
		--privileged \
		--mount type=volume,src=vol-dind,dst=/var/lib/docker \
		--name dind \
		docker:dind

exec:
	docker container exec \
		-it \
		dind \
		sh -c "docker container run -it --rm busybox sh -c 'ls -l'"

stop:
	docker container stop dind
	docker volume rm vol-dind
