IMAGE_NAME=try-docker/mysql
IMAGE_TAG=latest

default: run


clean:
	docker container stop mysql001
	docker network rm net1

build:
	docker network create net1
	docker image build -f Dockerfile -t $(IMAGE_NAME):$(IMAGE_TAG) .

run: build
	docker container run \
		--name mysql001 \
		--detach \
		--interactive \
		--tty \
		--rm \
		--net net1 \
		--publish 3306:3306 \
		-e MYSQL_ROOT_PASSWORD=rootroot \
		-e MYSQL_DATABASE=db1 \
		-e MYSQL_USER=user1 \
		-e MYSQL_PASSWORD=useruser \
		$(IMAGE_NAME):$(IMAGE_TAG) \
		--character-set-server=utf8mb4 \
		--collation-server=utf8mb4_unicode_ci \
		--default-authentication-plugin=mysql_native_password

cli:
# 繋がらない場合は一度コマンドラインで手動で以下のコマンドを走らせるとうまくいくときがある
	docker container exec --interactive --tty mysql001 mysql --user=root --password=rootroot --database=db1
