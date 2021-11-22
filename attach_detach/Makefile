start:
	docker container run \
		-it \
		--rm \
		--name attach-detach \
		--detach-keys="ctrl-\\" \
		busybox \
		sh

attach:
	docker container attach \
		--detach-keys="ctrl-\\" \
		attach-detach

stop:
	docker container kill attach-detach