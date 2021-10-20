default: run

clean:
	docker container stop reg001
	docker container rm reg001

build:
	docker image build -f Dockerfile -t try-docker/private_registry:latest .

run: build
# プライベートレジストリコンテナを起動
# レジストリも単なるコンテナとして利用できる
	@echo '-------------------------------------------------------------'
	docker container run \
		--name reg001 \
		--detach \
		--publish 5000:5000 \
		registry

# プライベートレジストリにアップするためにイメージにタグを付与
# 「レジストリの場所/リポジトリ名:バージョン番号」
	@echo '-------------------------------------------------------------'
	docker image tag try-docker/private_registry:latest localhost:5000/prireg:1.0

# docker image push でレジストリにイメージをアップ
	@echo '-------------------------------------------------------------'
	docker image push localhost:5000/prireg:1.0

# 結果確認のために一旦イメージを消す
	@echo '-------------------------------------------------------------'
	docker image rm try-docker/private_registry:latest localhost:5000/prireg:1.0

# イメージをpull
	@echo '-------------------------------------------------------------'
	docker image pull localhost:5000/prireg:1.0

# イメージがちゃんと存在するか確認
	@echo '-------------------------------------------------------------'
	docker image list
