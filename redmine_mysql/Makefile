default: run

clean:
	docker container stop redmine001 mysql001
	docker container rm redmine001 mysql001
	docker network rm net1

build:
	docker network create net1
	docker image build -f Dockerfile.mysql -t try-docker/mysql:latest .
	docker image build -f Dockerfile.redmine -t try-docker/redmine:latest .

run: build
# RedmineはMYSQLに接続できないとエラーとなり、コンテナが即死する.
# その場合は 何度か docker container restart redmine001 を試す
	docker container run --name mysql001 --detach --interactive --tty --network net1 try-docker/mysql:latest
	sleep 30
	docker container run --name redmine001 --detach  --interactive --tty --network net1 --publish 8086:3000 try-docker/redmine:latest
 

