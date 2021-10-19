default: run

clean:
	docker container stop wp001
	docker container stop mysql001
	docker container rm wp001
	docker container rm mysql001
	docker network rm net1

run: build run-mysql run-wp

build:
	docker network create net1

run-mysql:
	docker container run \
		--name mysql001 \
		--detach \
		--interactive \
		--tty \
		--net net1 \
		-e MYSQL_ROOT_PASSWORD=rootroot \
		-e MYSQL_DATABASE=db1 \
		-e MYSQL_USER=user1 \
		-e MYSQL_PASSWORD=useruser \
		mysql:latest \
		--character-set-server=utf8mb4 \
		--collation-server=utf8mb4_unicode_ci \
		--default-authentication-plugin=mysql_native_password

run-wp:
	docker container run \
		--name wp001 \
		--detach \
		--interactive \
		--tty \
		--net net1 \
		--publish 8085:80 \
		-e WORDPRESS_DB_HOST=mysql001 \
		-e WORDPRESS_DB_NAME=db1 \
		-e WORDPRESS_DB_USER=user1 \
		-e WORDPRESS_DB_PASSWORD=useruser \
		wordpress:latest