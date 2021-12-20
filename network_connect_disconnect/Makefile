CONTAINER_NAME := container-01
NETWORK_NAME   := network-conn-disconn-01
NETWORK_SUBNET := 192.168.149.0/24
NETWORK_FILTER := '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'

run: \
	_create_network \
	_start_container \
	_inspect_container_01 \
	_disconnect_from_bridge_network \
	_connect_another_network \
	_inspect_container_02 \
	_cleanup

_create_network:
	@echo '>>> Create docker network'
	docker network create --subnet $(NETWORK_SUBNET) $(NETWORK_NAME)
	@echo ''

_start_container:
	@echo '>>> Start container within docker host network'
	docker container run -dit --init --rm --name $(CONTAINER_NAME) --stop-signal SIGKILL busybox
	@echo ''

_inspect_container_01:
	@echo '>>> Inspect container network'
	docker container inspect --format=$(NETWORK_FILTER) $(CONTAINER_NAME)
	@echo ''

_disconnect_from_bridge_network:
	@echo '>>> Disconnect container from docker network'
	docker network disconnect bridge $(CONTAINER_NAME)
	@echo ''

_connect_another_network:
	@echo '>>> Connect container to docker network'
	docker network connect $(NETWORK_NAME) $(CONTAINER_NAME)
	@echo ''

_inspect_container_02:
	@echo '>>> Inspect container network'
	docker container inspect --format=$(NETWORK_FILTER) $(CONTAINER_NAME)
	@echo ''

_cleanup:
	@echo '>>> Cleanup'
	docker container stop $(CONTAINER_NAME) 2>&1 > /dev/null
	docker network rm $(NETWORK_NAME) 2>&1 > /dev/null
	@echo ''
