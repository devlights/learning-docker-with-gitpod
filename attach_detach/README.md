# これは何？

コンテナに対してのアタッチとデタッチのサンプルです。

# 実行

```sh
$ make start
$ make attach
$ make stop
```

```--detach-keys="ctrl-\\"``` として設定している

# デタッチするためのキーのデフォルト

```ctrl-p ctrl-q``` 

# デタッチキーを設定するオプション

```--detach-keys="ctrl-\\"``` のように設定する

```docker container run``` のとき、```docker container attach```のときに指定する。

---

Github Codespaces や Gitpod では、```ctrl-p``` にすでに役割が割り当てられているため
デフォルトのデタッチキーでは反応してくれない。なので、実行時に ```--detach-keys``` オプションで指定している。
