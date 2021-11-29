# 概要

コンテナに割り当てられたIPアドレスを知るには

```sh
$ docker container inspect コンテナID
```

とする。大量に出力されるが、その中の ```NetworkSettings.IPAddress``` が対象の項目。

```sh
$ docker container inspect --format="{{.NetrowkSettings.IPAddress}}" container-name
```
