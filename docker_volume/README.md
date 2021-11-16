# Docker の ボリュームについて

Docker には、３つのストレージが存在する

## volume

Dockerの管理化でストレージを管理する。

Linuxの場合は、 ```/var/lib/docker/volumes``` の下に配置される。

コンテナからじゃないと読み出せないが、OSのパスの違いなどが関係なくなったり

名前さえわかっていれば、他のコンテナからも扱えたりと便利。

バックアップはちょっとめんどくさい。

## bind

ホストのディレクトリ・ファイルを直接コンテナ側に共有する方法。

開発環境などを作る場合は、これを使うことがとても多い。

## tmpfs

名前の通り、一時的なストレージとして使うもの。メモリにマッピングされる。

# ```-v (--volume)``` オプションと ```--mount``` オプション

公式では、```--mount``` の方を推奨している模様。

```-v``` オプションの場合は、３つのストレージ方式のどれをつかうのかが明確ではないが

```--mount``` オプションの場合は、分かりやすい。ただし、ちょっと長くなる。

```sh
$ docker container run -it --mount type=volume,src=vol-1,dst=/app debian:bullseye bash
$ docker container run -it --mount type=bind,src=${PWD},dst=/app debian:bullseye bash
```

# 参考情報

- [Docker の Volume がよくわからないから調べた](https://qiita.com/aki_55p/items/63c47214cab7bcb027e0)
- [Docker、ボリューム(Volume)について真面目に調べた](https://qiita.com/gounx2/items/23b0dc8b8b95cc629f32)
