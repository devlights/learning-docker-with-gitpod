default: container_attach

#
# docker container run -it --rm --name lifecycle01 busybox sh と 以下は同じ
#

image_pull:
	docker image pull busybox

container_create: image_pull
	docker container create -it --rm --name lifecyble01 busybox sh

container_start: container_create
	docker container start lifecyble01

container_attach: container_start
	docker container attach lifecyble01
