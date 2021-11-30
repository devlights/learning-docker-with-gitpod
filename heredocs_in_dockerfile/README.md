# Dockerfile で ヒアドキュメント を使用

使用するためには、以下の条件がある。

- ```DOCKER_BUILDKIT``` が有効になっている
- Dockerfileのシンタックスバージョンが ```docker/dockerfile:1.3-labs``` 以降

```DOCKER_BUILDKIT=1``` を毎回指定しても良いが、```docker bulidx build``` でイメージをビルドしても良い。

Dockerfileの先頭行に以下のようにシンタックスバージョンを指定する必要がある。

```Dockerfile
# syntax docker/dockerfile:1-labs
FROM golang:1.17-alpine

.
.
.
```

# 例

```sh
$ DOCKER_BUILDKIT=1 docker image build ...
```

```sh
$ docker buildx build ...
```

# 参考情報

- [Dockerfile で新しく使えるようになった構文「ヒアドキュメント」で複数行の RUN をシュッと書く](https://kakakakakku.hatenablog.com/entry/2021/08/10/085625)
- [Introduction to heredocs in Dockerfiles](https://www.docker.com/blog/introduction-to-heredocs-in-dockerfiles/)
