# try-docker
This is a repository for learning Docker using Gitpod. (for myself)

# Launch docker daemon in Gitpod

基本、最初からDockerデーモンは起動しているので以下のコマンドは実施する必要はない。

```sh
$ sudo docker-up
```

- https://www.gitpod.io/blog/root-docker-and-vscode


# WSL2

WSL2が出るまでは、WindowsのHomeエディションではデスクトップ版のDockerが利用できなかった。

2020年春頃のWindowsアップデートにて、WSL2がリリースされたため、現在はHomeエディションでもデスクトップ版Dockerが利用できるようになった。

# メモ

コンテナ操作の基本は、コマンドでの命令。

コンテナを操作するコマンド文はすべて「docker」から始まる。
dockerと書いてから、「何を」「どうする」「対象」などを記述して実行する。

dockerコマンドに続けて書く「何を」「どうする」の部分を「コマンド」と呼ぶ。
コマンドは「上位コマンド」と「副コマンド」で構成されている。

上位コマンドの部分が「何を」、副コマンドの部分が「どうする」に該当する。
また、「対象」の部分には、コンテナ名やイメージ名などの具体的な名前が入る。

##### 例

```sh
$ docker コマンド 対象
```

```sh
$ docker image pull xxxxx
$ docker container start xxxxx
```

```sh
$ docker container run -d xxxxx                            # -d は「バックグラウンドで実行する」の意味
$ docker container run --interactive --tty --detach xxxxx  # docker run -it -d xxxxx と同じ
```

歴史的経緯から、「start」や「run」のように ```container``` を付けなくても実行できるコマンドがある。
Docker 1.13にて、コマンドの再編成が実施され、上位コマンドと副コマンドの組み合わせに統一された。
旧形式の実行方法も現在のところ可能。

##### 旧形式

```sh
$ docker run xxxxx
```

##### 新形式

```sh
$ docker container run xxxxx
```

ヘルプを見たい場合は以下のようにする。

```sh
$ docker --help
$ docker image --help
$ docker container run --help 
```

上記は、それぞれ出力されるヘルプの内容が異なる。

インストールされているDockerの情報を調べたい場合は ```docker version``` とする。

```sh
gitpod /workspace/try-docker $ docker version
Client: Docker Engine - Community
 Version:           20.10.8
 API version:       1.41
 Go version:        go1.16.6
 Git commit:        3967b7d
 Built:             Fri Jul 30 19:54:27 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.8
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.16.6
  Git commit:       75249d8
  Built:            Fri Jul 30 19:52:33 2021
  OS/Arch:          linux/amd64
  Experimental:     true
 containerd:
  Version:          1.4.9
  GitCommit:        e25210fe30a0a703442421b0f60afac609f950a3
 gitpod:
  Version:          1.0.1
  GitCommit:        v1.0.1-0-g4144b63
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```

```docker run``` は、以下のコマンドの動作をひとまとめにしたもの。

```sh
$ docker image pull
$ docker container create
$ docker container start
```

ホストOSの特定のディレクトリをマウントした状態で起動したい場合は以下のようにする。

```sh
$ docker run -it -v ${PWD}:/app -w container-working-directory --name container-name image-name
```

```sh
$ docker run -it -v ${PWD}:/app -w /app --name mycontainer alpine:latest
```

```docker ps``` コマンドと ```docker container list``` は同じ。

```sh
$ docker ps [-a]
$ docker container list [-a]
```

```docker ps``` が古い形式。



