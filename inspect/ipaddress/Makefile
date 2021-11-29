run:
	docker container run \
		--init \
		--detach \
		--interactive \
		--rm \
		--name inspect-ipaddr \
		busybox
	docker container inspect \
		--format="{{.NetworkSettings.IPAddress}}" \
		inspect-ipaddr
	docker container stop inspect-ipaddr
