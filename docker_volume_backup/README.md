# ボリュームのバックアップの考え方

Dockerでボリュームをバックアップするときは、適当なコンテナを割り当てて、そのコンテナを使ってバックアップを取る。

例を上げると、適当なLinuxシステムが入ったコンテナを一つ起動し、そのコンテナの /tmp などのディレクトリにバックアップ対象の
コンテナをバインドマウントして、tarでバックアップを取る。そのバックアップをDockerホスト側で取り出せば、バックアップは完了となる。

リストアするときは、この逆を行う。コンテナを一つ起動し、そのコンテナの /tmp などのディレクトリにレストア対象のDockerボリュームを
ボリュームマウントする。コンテナ内からバックアップファイルを解凍して、マウントポイントに流し込めばレストアが完了する。

なお、Dockerボリュームは複数のコンテナから同時にマウント可能。バックアップを取る際に利用中のコンテナのマウントは外す必要は特に無い。
ただし、コンテナが稼働中であると、そのコンテナが書き込みを行ったりしてしまい、データの整合性が取れなくなる可能性があるので、できれば停止しておいた方が無難。

# 実行例

```sh
$ make
+----------------------------------------------+
+ バックアップ対象 Docker ボリューム作成
+----------------------------------------------+
docker volume create vol-backup-target
vol-backup-target
+----------------------------------------------+
+ バックアップ対象 Docker ボリューム データ書き込み
+----------------------------------------------+
docker container run \
        --init \
        --rm \
        --mount type=volume,src=vol-backup-target,dst=/data \
        --workdir /data \
        busybox \
        sh -c 'pwd ; echo "test data" > test.txt'
/data
+----------------------------------------------+
+ バックアップ対象 Docker ボリューム データ確認
+----------------------------------------------+
docker container run \
        --init \
        --rm \
        --mount type=volume,src=vol-backup-target,dst=/data \
        --workdir /data \
        busybox \
        sh -c 'pwd ; ls -lh'
/data
total 4K     
-rw-r--r--    1 root     root          10 Nov 26 07:57 test.txt
+----------------------------------------------+
+ バックアップ対象 Docker ボリューム バックアップ
+----------------------------------------------+
sudo rm -rf ./backup.tar.gz
docker container run \
        --init \
        --rm \
        --mount type=volume,src=vol-backup-target,dst=/data \
        --mount type=bind,src=/workspaces/try-docker/docker_volume_backup,dst=/backup \
        --workdir / \
        busybox \
        sh -c 'tar czf /backup/backup.tar.gz /data'
tar: removing leading '/' from member names
+----------------------------------------------+
+ バックアップファイルの中身を確認
+----------------------------------------------+
sudo tar ztf /workspaces/try-docker/docker_volume_backup/backup.tar.gz
data/
data/test.txt
+----------------------------------------------+
+ バックアップ対象 Docker ボリューム削除
+----------------------------------------------+
docker volume rm vol-backup-target
vol-backup-target
+----------------------------------------------+
+ レストア対象 Docker ボリューム作成
+----------------------------------------------+
docker volume create vol-restore-target
vol-restore-target
+----------------------------------------------+
+ レストア対象 Docker ボリューム レストア
+----------------------------------------------+
docker container run \
        --init \
        --rm \
        --mount type=volume,src=vol-restore-target,dst=/data \
        --mount type=bind,src=/workspaces/try-docker/docker_volume_backup,dst=/backup \
        --workdir / \
        busybox \
        sh -c 'tar zxf /backup/backup.tar.gz'
+----------------------------------------------+
+ レストア対象 Docker ボリューム データ確認
+----------------------------------------------+
docker container run \
        --init \
        --rm \
        --mount type=volume,src=vol-restore-target,dst=/data \
        --workdir /data \
        busybox \
        sh -c 'pwd ; ls -lh'
/data
total 4K     
-rw-r--r--    1 root     root          10 Nov 26 07:57 test.txt
+----------------------------------------------+
+ レストア対象 Docker ボリューム削除
+----------------------------------------------+
docker volume rm vol-restore-target
vol-restore-target
```