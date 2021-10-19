default: run

clean:
	docker container stop mount001
	docker container rm mount001

build:
	docker image build -f Dockerfile -t try-docker/bind_mount:latest .

run: build
	docker container run \
		--name mount001 \
		-dit \
		-v "${PWD}/..":/app \
		-v "${PWD}":/app2 \
		try-docker/bind_mount:latest
	
	@echo '------------------------------------------------------------------------'
	docker container exec -it mount001 bash -c "ls -l /app"
	@echo '------------------------------------------------------------------------'
	docker container exec -it mount001 bash -c "ls -l /app2"