default: run

clean:
	docker-compose down

build:
	docker-compose build

run: build
	docker-compose up -d
	@echo '------------------------------------------------------------------------'
	docker container exec -it mount001 bash -c "ls -l /app"
	@echo '------------------------------------------------------------------------'
	docker container exec -it mount001 bash -c "ls -l /app2"