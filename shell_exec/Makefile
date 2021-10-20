default: run

clean:
	docker container stop shellexec001
	docker container rm shellexec001

build:
	docker image build -f Dockerfile -t try-docker/shell_exec:latest .

run: build
	docker container run -dit --name shellexec001 try-docker/shell_exec:latest
	docker container exec -it shellexec001 /bin/bash -c 'bash --version; sleep 3; exit'
	docker container exec -it shellexec001 /bin/bash