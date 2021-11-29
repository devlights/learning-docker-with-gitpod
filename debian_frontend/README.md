# 概要

Debian 系Linuxにて、aptを使ってなにかをインストールしようとすると、たまにインタラクティブなインストーラーが
あったりして、処理が止まってしまう場合がある。

それを防ぐために

```sh
$ apt-get update -q && DEBIAN_FRONTEND=noninteractive apt-get install -yq xxxx
```

とすることが多い。

また、同じように ```--no-install-recommends``` オプションをつけることも多い。

# 参考情報

- http://docs.docker.jp/v1.11/engine/faq.html#dockerfile-debian-frontend-noninteractive
