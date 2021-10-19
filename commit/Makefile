default: run

clean:
	docker container stop commit001 commit002
	docker container rm commit001 commit002
	docker image rm try-docker/commit1 try-docker/commit2

build:
	docker image build -f Dockerfile -t try-docker/commit1:latest .

run: build
	docker container run --name commit001 -dit try-docker/commit1:latest
	docker container exec -it commit001 apk add bash vim
	docker container commit commit001 try-docker/commit2:latest
	docker container run --name commit002 -dit try-docker/commit2:latest
	docker container exec -it commit002 bash -c "vim --version"
	@echo '-----------------------------------------------------------------'
	docker image list