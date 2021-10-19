default: run

clean:
	docker container stop filecp001
	docker container rm filecp001

build:
	docker image build -f Dockerfile -t try-docker/file_copy:latest .

run: build
	docker container run --name filecp001 -dit try-docker/file_copy:latest	
	@echo '[HOST ----> CONTAINER]'
	docker container cp ../README.md filecp001:/app/HOST_TO_CONTAINER
	@echo '-------------------------------------------------------'
	docker container exec filecp001 ls -l /app
	@echo '-------------------------------------------------------'
	@echo '[CONTAINER ----> HOST]'
	docker container cp filecp001:/app/HOST_TO_CONTAINER ./CONTAINER_TO_HOST
	@echo '-------------------------------------------------------'
	ls -l
