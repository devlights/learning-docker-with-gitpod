all: make-volume exec rm-volume

make-volume:
	docker volume create vol-1
	docker volume create vol-2

rm-volume:
	docker volume rm vol-1
	docker volume rm vol-2

exec:
	docker image pull debian:bullseye
	@echo '------- vol-1 に helloworld.txt を作成 -------'
	docker container run \
		--rm \
		--mount type=volume,src=vol-1,dst=/app \
		--workdir /app \
		debian:bullseye \
		bash -c 'echo "helloworld" > helloworld.txt'
	@echo '------- 別のコンテナでvol-1をマウントして確認 -------'
	docker container run \
		--rm \
		--mount type=volume,src=vol-1,dst=/app \
		--workdir /app \
		debian:bullseye \
		bash -c 'ls -l; cat helloworld.txt'	
	@echo '------- vol-2 に worldhello.txt を作成 -------'
	docker container run \
		--rm \
		--mount type=volume,src=vol-2,dst=/app \
		--workdir /app \
		debian:bullseye \
		bash -c 'echo "worldhello" > worldhello.txt'

# 同じマウントポイントに複数の --mount を指定するとエラーとなる
#   docker: Error response from daemon: Duplicate mount point: /app.
#
#	docker container run \
#		--rm \
#		--mount type=volume,src=vol-1,dst=/app \
#		--mount type=volume,src=vol-2,dst=/app \
#		debian:bullseye \
#		bash -c 'ls -al'

	@echo '------- vol-1,vol-2 をマウント -------'
	docker container run \
		--rm \
		--mount type=volume,src=vol-1,dst=/app1 \
		--mount type=volume,src=vol-2,dst=/app2 \
		debian:bullseye \
		bash -c 'ls -l /app1 /app2'
	@echo '------- ホストディレクトリをバインドマウント -------'
	docker container run \
		--rm \
		--mount type=bind,src=${PWD},dst=/app \
		--workdir /app \
		debian:bullseye \
		bash -c 'ls -l /app'
