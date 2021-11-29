# daemon.json とは

```daemon.json``` は、Dockerエンジンの各種オプションを指定することができる設定ファイル。

このファイルは、 Dockerのページに記載されている ```apt``` などのパッケージ管理システムを用いて

インストールした場合は、デフォルトで「存在しない」。なので、自分でファイルを追加する必要がある。

ファイルの場所は ```/etc/docker/daemon.json``` 。

ここにオプションを記載すると、Dockerエンジン起動時にデフォルトで読み込まれる。

公式ドキュメントにも以下の記載があり、```daemon.json``` を利用することが推奨されている。

> The recommended way is to use the platform-independent daemon.json file, which is located in /etc/docker/ on Linux by default. 

# 試した環境

chromebook内のLinux環境にDockerをインストールしている状態。

```sh
$ cat /etc/os-release 
PRETTY_NAME="Debian GNU/Linux 10 (buster)"
NAME="Debian GNU/Linux"
VERSION_ID="10"
VERSION="10 (buster)"
VERSION_CODENAME=buster
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"

$ sudo docker version
Client: Docker Engine - Community
 Version:           20.10.11
 API version:       1.41
 Go version:        go1.16.9
 Git commit:        dea9396
 Built:             Thu Nov 18 00:36:10 2021
 OS/Arch:           linux/arm64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.11
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.16.9
  Git commit:       847da18
  Built:            Thu Nov 18 00:34:45 2021
  OS/Arch:          linux/arm64
  Experimental:     false
 containerd:
  Version:          1.4.12
  GitCommit:        7b11cfaabd73bb80907dd23182b9347b4245eb5d
 runc:
  Version:          1.0.2
  GitCommit:        v1.0.2-0-g52b36a2
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```

# insecure-registries オプションを指定してみる

insecure-registries オプションは

プライベートレジストリを構築した際、デフォルトでは docker エンジンは https 通信しか許可していないが

社内などのクローズドな環境で利用する場合は http 通信を許可したい場合などに設定するオプション。

以下、 ```/etc/docker/daemon.json``` が存在していないものとして話を進める。

まず、現在の状況を確認。

```sh
$ sudo docker info 2> /dev/null | tail -5
 Experimental: false
 Insecure Registries:
  127.0.0.0/8
 Live Restore Enabled: false
```

```Insecure Registries:``` は、```127.0.0.0/8``` だけとなっている。

ここに、たとえば ```192.168.1.10:5000``` を追加したい場合、```/etc/docker/daemon.json``` に設定を追加する。

```sh
$ sudo sh -c 'cat << EOF > /etc/docker/daemon.json
> {
>       "insecure-registries":[ "192.168.1.10:5000" ]
> }
> EOF
> '
```

追加後、デーモンを再起動する。

```sh
$ sudo systemctl restart docker
```

設定が読み込まれたか確認。

```sh
$ sudo docker info 2> /dev/null | tail -5
 Insecure Registries:
  192.168.1.10:5000
  127.0.0.0/8
 Live Restore Enabled: false
```

ちゃんと読み込まれている。

# 参考情報

- [Custom Docker daemon options](https://docs.docker.com/config/daemon/systemd/#custom-docker-daemon-options)
- [Configure the Docker daemon](https://docs.docker.com/config/daemon/#configure-the-docker-daemon)
