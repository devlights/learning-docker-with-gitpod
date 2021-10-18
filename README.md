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