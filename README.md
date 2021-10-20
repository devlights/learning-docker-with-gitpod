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

コンテナIDやイメージIDを指定する際は、一意な部分のみを指定できれば良い。

なので、例えば

```sh
$ docker image list -a
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
httpd        latest    d54056386fbb   5 days ago    138MB
alpine       latest    14119a10abf4   7 weeks ago   5.59MB
```

となっていたら

```sh
$ docker image rm d5
```

コンテナを破棄するためには、停止する必要がある。コンテナは動いているものをいきなり削除できない。

```sh
$ docker container stop xxxxx
$ docker container rm xxxxx
```

コンテナはデフォルトでは外から通信アクセスできない状態となっている。アクセスできるように設定するためには、コンテナ作成時に設定する必要がある。

基本的に作成後に変更は出来ない。ポートを指定する場合は ```-p 8888:80``` のように指定する。前がホスト側のポート番号、後ろがコンテナ側のポート番号。

```sh
$ docker run -d --name apa001 -p 8888:80 httpd:latest
```

dockerコマンド実行時に環境変数を設定する場合は ```-e``` オプションを指定する。
環境変数一つずつに ```-e``` が必要となる。例えば、３つ指定する場合は

```sh
-e XXX=XXX -e XXX=XXX -e XXX=XXX
```

となる。

コンテナを削除しても、イメージは残ったままとなる。イメージを削除するためには、そのイメージから作成したコンテナが全て削除されていないといけない。

コンテナ同士を繋ぐ場合、ただ普通にコンテナを作っただけでは、コンテナは繋がらないので、仮想的なネットワークを作り、そこに両方のコンテナを所属させることで、コンテナ同士を繋げる。

仮想的なネットワークを作るコマンドが ```docker network create``` となる。

```sh
$ docker network create xxx
$ docker network rm xxx
```

dockerコマンド実行時にネットワークを設置する場合は ```--net``` オプションを指定する。

# コンテナとホスト間でファイルをコピーする

コピーは、コンテナからホスト、ホストからコンテナのどちらも可能。

```docker container cp``` コマンドを使う

## ホストからコンテナ

```sh
$ docker container cp ../README.md filecp001:/app
```

## コンテナからホスト

```sh
$ docker container cp filecp001:/app/README.md ./README2.md
```

# ボリュームのマウント

実際にコンテナを使うのであれば、記憶領域のマウントが必須。なぜなら、そこにデータを置くから。

記憶領域のマウント方法は2種類ある。

- ボリュームマウント
  - Dockerエンジンが管理している領域内にボリュームを作成し、ディスクとしてコンテナにマウントする。
  - 名前だけで管理できるので手軽に扱える反面、ボリュームに対して直接操作しづらい。
    - 仮で使いたい場合
    - 滅多に触らないけど、消してはいけないファイル
- バインドマウント
  - Dockerをインストールしたパソコンの場所、つまり、Dockerエンジンが管理していない場所の既に存在するディレクトリをコンテナにマウントする。
  - ディレクトリだけではなく、ファイル単位でもマウントできる。
  - ディレクトリに対して直接ファイルを置いたり開いたりできるので、頻繁に触るファイルはここに置くべき。

バインドマウントは、Dockerの管理外の好きな場所におけるので、普通のファイルと同じように扱えるのが売り。

この他にも「一時メモリ(tmpfs)マウント」というのもある。その名前の通り、メモリにマウントする方法。当然コンテナの再起動で消える。

ボリュームマウント、バインドマウントのどちらの方法でも、記憶領域のマウントは ```docker container run``` コマンドにオプションとして指定する。

```sh
$ docker container run -v ホストのパス:コンテナのパス
$ docker container run --volume ホストのパス:コンテナのパス
```

マウントしたい場所が複数存在する場合は、```-v (--volume)``` オプションを複数回指定すれば良い。

# コンテナのイメージ化

コンテナを別の環境にコピーしたい場合などに利用するので、サーバエンジニアは必須の知識となる。

イメージの作成方法には２つある。

- すでにあるコンテナを ```docker container commit``` でイメージの書き出しをする方法
- ```Dockerfile``` でイメージを作る方法

## commit で書き出す方法

```sh
$ docker container commit commit001 try-docker/commit2:latest
```

## Dockerfileを使う方法

```sh
$ docker build -f Dockerfileのパス -t イメージ名:タグ 材料フォルダのパス
```

通常は以下のようになる。

```sh
$ docker build -f Dockerfile -t xxx:1.0 .
```

Dockerfile の中で使う命令は以下のようなものとなる。

- FROM
  - 元にするイメージを指定する
- ADD
  - イメージにファイルやフォルダを追加する
- COPY
  - イメージにファイルやフォルダを追加する
- RUN
  - イメージをビルドするときにコマンドを実行する
- CMD
  - コンテナを起動するときに実行する既存のコマンドを指定する
- ENTRYPOINT
  - イメージを実行するときのコマンドを強要する
- ONBUILD
  - ビルド完了したときに任意の命令を実行する
- EXPOSE
  - 通信を想定するポートをイメージの利用者に伝える
- VOLUME
  - 永続データが保存される場所をイメージの利用者に伝える
- ENV
  - 環境変数を定義する
- WORKDIR
  - RUN, CMD, ENTRYPOINT, ADD, COPYの際の作業ディレクトリを指定する
- SHELL
  - ビルド時のシェルを指定する
- LABEL
  - 名前やバージョン番号、製作者情報などを設定する
- USER
  - RUN, CMD, ENTRYPOINTで指定するコマンドを実行するユーザやグループを設定する
- ARG
  - ```docker image build```する際に指定できる引数を宣言する
- STOPSIGNAL
  - ```docker container stop``` する際に、コンテナで実行しているプログラムに対して送信するシグナルを変更する
- HEALTHCHECK
  - コンテナの死活確認をするヘルスチェックの方法をカスタマイズする


# イメージのSaveとLoad（持ち運び）

コンテナはそのままでは移動・コピーが出来ない。一旦イメージにする必要がある。

ただし、イメージもそのままでは利用できないので、Dockerレジストリを経由させるか、 ```docker image save``` コマンドで tar ファイルにする。

ファイルからイメージとして取り込みたい場合は ```docker image load``` コマンドを利用する。

```sh
$ docker image save -o xxxx.tar イメージ名
```

```sh
$ docker image load --input xxxx.tar
```

# コンテナの改造

## シェルの起動

コンテナは、何も指定せずに起動すると、当然シェルも動いていない状態となる。

なので、bashなどのシェルを起動して命令を受け取ってもらうようにする必要がある。

起動中のコンテナに対して、シェル起動して中に入るには ```docker container exec``` を使う。

```sh
$ docker container exec -it container-name /bin/bash
```

ようにする。

```docker container run``` に対して引数でシェルを指定することもできるが、その場合コンテナに入っているソフトウェアを動かす代わりに
シェルを動かすことになる。なので、コンテナは作られているものの、ソフトウェアがスタートしていない状態となる。シェルでの操作が終わった後に
改めて ```docker container start``` でスタートさせる必要が出てくる。

なので、```docker container exec``` を使う。


# プライベートレジストリ

普通に ```docker image pull``` や ```docker container run``` を利用すると デフォルト では [DockerHub](https://hub.docker.com/) からイメージをダウンロードする。

イメージの配布場所を ```Dockerレジストリ``` と言う。一般に公開されているか否かに関わらず、配布場所と言えばDockerレジストリと呼ぶ。

[DockerHub](https://hub.docker.com/)は、Dockerレジストリのうち、Docker社の公式が運営しているもの。

当然、プライベートなDockerレジストリも作成することができる。

## レジストリとリポジトリ

レジストリ（登記所）とリポジトリ（倉庫）は似ているが違う。

レジストリは ```イメージの配布場所``` 。一方、リポジトリは、レジストリの中をさらに区切った単位。

一つのレジストリの中にNのリポジトリが存在する。

## アップロード方法

イメージのアップロード先が [DockerHub](https://hub.docker.com/) であっても、プライベートなDockerレジストリであってもイメージにはタグをつける必要がある。

タグ名は以下のような命名となる。

```レジストリの場所/リポジトリ名:バージョン番号```

例として以下のようになる

#### localhost:5000 で公開されているレジストリで、priregというリポジトリ名で、バージョン番号が1.0

```localhost:5000/prireg:1.0```

(*) [DockerHub](https://hub.docker.com/) の場合は ```DockerHubのID/リポジトリ名:バージョン番号``` となる

イメージにタグを付与するには ```docker image tag``` コマンドを使う

```sh
$ docker image tag 元のイメージ名 レジストリの場所/リポジトリ名:バージョン番号
```

```sh
$ docker image tag try-docker/prireg:latest localhost:5000/prireg:1.0
```

イメージをアップするには ```docker image push``` コマンドを使う

```sh
$ docker image push レジストリの場所/リポジトリ名:バージョン番号
```

```sh
$ docker image push localhost:5000/prireg:1.0
```

## プライベートレジストリを作る

Dockerでは、レジストリも一つのコンテナとして表現される。

レジストリ用のイメージが提供されているので、以下のようにコンテナを起動すると、そのホストの中にプライベートレジストリが起動する。

```sh
$ docker container run --name reg001 -d -p 5000:5000 registry 
```

# Docker Compose

Dockerでの構築作業に関わるコマンド文の内容を１つのテキストファイル（定義ファイル)(YAML) に書き込んで、一気に実行したり停止・破棄したりするのが ```Docker Compose``` 。

実際にDockerで作業する場合は、dockerコマンドを手打ちで頑張ることはあまりなくて、このDocker Composeを使って作業することが多い。

ファイルの名前は通常 ```docker-compose.yml``` とする。

起動は ```docker-compose -f 定義ファイル up オプション``` とする。

停止は ```docker-compose -f 定義ファイル down オプション``` とする。

定義ファイルには、コンテナやボリュームを「こういう設定で作りたい」という項目を書いておく。

Dockerfileと似ているが、Dockerfileはイメージを作るものなので、ネットワークやボリュームなどは作れない。

Kubernetesとも似ている感じがするが、KubernetesはDockerコンテナを管理するもので、Docker Composeはコンテナなどを作って消すだけで管理機能は持っていない。

Docker Compose は、WindowsとMacの場合はDockerデスクトップ版に付属しているので追加のインストール作業は必要ない。

Linuxの場合は Docker Compose と Python3 のインストールが必要となる。

インストール方法は [Docker Composeのインストール](https://docs.docker.jp/compose/install.html#linux) を参照。

docker-composeのバージョンは以下で確認できる。

```sh
$ docker-compose version
docker-compose version 1.29.2, build 5becea4c
docker-py version: 5.0.0
CPython version: 3.7.10
OpenSSL version: OpenSSL 1.1.0l  10 Sep 2019
```

docker-compose.yml の内容は例えば以下のようになる。

```yaml
version: "3"

services:
  httpd001:
    build:
      context: .
      dockerfile: Dockerfile.apache
    networks:
      - net1
    ports:
      - 8085:80
    restart: always
  tomcat001:
    build:
      context: .
      dockerfile: Dockerfile.tomcat
    networks:
      - net1
    restart: always
    depends_on:
      - httpd001
networks:
  net1:

```

後は、docker-compose.yml が存在するディレクトリに移動して

```sh
$ docker-compose up -d
```

で起動して

```sh
$ docker-compose down
```

で停止となる。


# Kubernetes

Kubernetesは複数の物理的マシンに複数のコンテナがあることが前提。

Dockerは、1代のブルリ的マシンで実行するイメージだが、Kubernetesは複数の物理的マシンがあることが前提。

さらに、その1台の中に複数のコンテナがあるという状況で利用価値がある。

複数のコンテナを1台ずつ作ったり、管理したりするのは大変なので、Kubernetesが肩代わりして管理してくれるイメージ。

Kubernetesは、このようなコンテナの作成や管理の煩雑さをうまくやってくれるツールで、Docker Composeのような定義ファイル（マニフェストファイル）を
つくっておけば、それに従って対象の物理マシンにコンテナを作成して管理してくれる。

## マスタノードとワーカーノード

Kubernetesは マスターノード と呼ばれるコントロールを司るノードと ワーカーノード と呼ばれる実際に動かすノードで構成される。

マスターノードは、現場監督のようなもの。マスターノード上でコンテナは動かない。ワーカーノード上のコンテナを管理するのが仕事。

なので、マスターノードにはDocker Engineもインストールしない。

ワーカーノードは、実際のサーバに当たる部分で、コンテナはここで動く。

このように、マスターノードとワーカーノードで構成されたKubernetesシステムの一群を クラスター と呼ぶ。


## Kubernetes 動作に必要なもの

- マスターノード
  - Kubernetes 本体
  - CNI (仮想ネットワークドライバ)
    - flannel
    - Clico
    - AWS VPC CNI
  - etcd
- ワーカーノード
  - Kubernetes 本体
  - CNI
  - Docker Engine

また、マスターノードを管理するために 管理者 のマシンには kubectl というものを入れる必要がある。


## Kubernetes は、状態を維持する

Kubernetes では、あるべき姿（コンテナをN個、ボリュームをN個で構成）をYAML形式のマニフェストファイルに記載して動かす。

Docker Composeは、オプションの指定によって手動でコンテナの数を返ることはできるが、監視はしていないので、作成時しか関与しない「作って終わり」のもの。

Kubernetesは、「その状態を維持」する。

なので、なんらかの理由でコンテナが壊れた場合、Kubernetesが勝手に壊れたコンテナを削除して、新しく作り直したりしてくれる。（ちゃんと設定していれば）


## 最低限知っておくべき用語

- Pod
  - コンテナとボリュームのセット
  - Kubernetesが管理する最小単位
- サービス
  - Podへのアクセスを管理
- デプロイメント
  - Podのデプロイを管理

同一構成のPodの塊を レプリカ (Replica) という。


## 最低限知っておくべき事実

Kubernetesを「ちゃんと」使うには、しっかりとして大量の知識とネットワークに関する深い知識などが必要になる。

さらに、Kubernetesは本来大規模なシステムが前提であり、マスターノードとワーカーノードは別々の物理的マシンにされているなど、簡単に手が出せるものではない。

ClusterIPの設定やその前に置くロードバランサなどはインフラ屋さんに頼まないといけない。

なので、現実的に大企業に属しているなどじゃないと、イチから構築などはまずチャンスがない。

ほとんどの人はクラウドを利用してKubernetesに触れることになる。

単に触ってみる程度であれば、Dockerデスクトップ版でお試しKubernetesがインストールできるので、それで触ってみるのがいいかもしれない。

Linuxの場合はMinikubeをインストールして操作してみるという感じになる。

WSL2環境のDebian/sidに Minikube をインストールしてみた感じ、以下のようになった。

```sh
$ cat /etc/os-release
cat /etc/os-release
PRETTY_NAME="Debian GNU/Linux bookworm/sid"
NAME="Debian GNU/Linux"
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"



$ curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
sudo dpkg -i minikube_latest_amd64.deb
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 22.2M  100 22.2M    0     0  21.9M      0  0:00:01  0:00:01 --:--:-- 21.9M
Selecting previously unselected package minikube.
(Reading database ... 39822 files and directories currently installed.)
Preparing to unpack minikube_latest_amd64.deb ...
Unpacking minikube (1.23.2-0) ...
Setting up minikube (1.23.2-0) ...


$ sudo dpkg -i minikube_latest_amd64.deb
(Reading database ... 39823 files and directories currently installed.)
Preparing to unpack minikube_latest_amd64.deb ...
Unpacking minikube (1.23.2-0) over (1.23.2-0) ...
Setting up minikube (1.23.2-0) ...


$ minikube start --vm-driver=none
😄  minikube v1.23.2 on Debian bookworm/sid
✨  Using the none driver based on user configuration

🤷  Exiting due to PROVIDER_NONE_NOT_FOUND: The 'none' provider was not found: exec: "iptables": executable file not found in $PATH
💡  Suggestion: iptables must be installed
📘  Documentation: https://minikube.sigs.k8s.io/docs/reference/drivers/none/
```

ちなみに Gitpod 上ではエラーが出て Minikube が動作しなかった。

```sh
$ curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube_latest_amd64.deb
$ sudo dpkg -i minikube_latest_amd64.deb
$ sudo apt update && sudo apt install conntrack

$ minikube start --vm-driver=none
😄  minikube v1.23.2 on Ubuntu 20.04 (amd64)
✨  Using the none driver based on user configuration
👍  Starting control plane node minikube in cluster minikube
🤹  Running on localhost (CPUs=16, Memory=64319MB, Disk=595282MB) ...
ℹ️  OS release is Ubuntu 20.04.2 LTS

❌  Exiting due to RUNTIME_ENABLE: sudo systemctl daemon-reload: exit status 1
stdout:

stderr:
System has not been booted with systemd as init system (PID 1). Can't operate.
Failed to connect to bus: Host is down


╭───────────────────────────────────────────────────────────────────────────────────────────╮
│                                                                                           │
│    😿  If the above advice does not help, please let us know:                             │
│    👉  https://github.com/kubernetes/minikube/issues/new/choose                           │
│                                                                                           │
│    Please run `minikube logs --file=logs.txt` and attach logs.txt to the GitHub issue.    │
│                                                                                           │
╰───────────────────────────────────────────────────────────────────────────────────────────╯
```

## Windows上のDockerデスクトップでKubernetesを有効にして試してみた例

#### デプロイメント

ファイル名は ```httpd001dep.yml``` とした

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: httpd001dep
spec:
  selector:
    matchLabels:
      app: httpd001kube
  replicas: 1
  template:
    metadata:
      labels:
        app: httpd001kube
    spec:
      containers:
        - name: httpd001
          image: httpd
          ports:
            - containerPort: 80
```

#### サービス

ファイル名は ```httpd001ser.yml``` とした

```yaml
apiVersion: v1
kind: Service
metadata:
  name: httpd001ser
spec:
  type: NodePort
  ports:
  - port: 8099
    targetPort: 80
    protocol: TCP
    nodePort: 30080
  selector:
    app: httpd001kube
```

#### 実行

デプロイメントを適用する

```sh
$ kubectl apply -f ./httpd001dep.yml
deployment.apps/httpd001dep created
```

Podの状況を見てみる。

```sh
$ kubectl get pods
NAME                          READY   STATUS    RESTARTS   AGE
httpd001dep-945d9f6dd-6q997   1/1     Running   0          14s
```

レプリカの数を１にしているので、起動しているコンテナは１つ。

次にレプリカの数を３に変更して再度適用。

```sh
$ kubectl apply -f ./httpd001dep.yml
deployment.apps/httpd001dep created
```

```sh
$ kubectl get pods
NAME                          READY   STATUS              RESTARTS   AGE
httpd001dep-945d9f6dd-2h6dz   0/1     ContainerCreating   0          1s
httpd001dep-945d9f6dd-6q997   1/1     Running             0          22s
httpd001dep-945d9f6dd-sz6gg   0/1     ContainerCreating   0          1s
```

コンテナの数が自動的に調整された。


サービスを適用する

```sh
$ kubectl apply -f ./httpd001ser.yml
service/httpd001ser created
```

サービスの状況を確認

```sh
$ kubectl get services
NAME          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
httpd001ser   NodePort    10.107.153.247   <none>        8099:30080/TCP   6s
kubernetes    ClusterIP   10.96.0.1        <none>        443/TCP          65m
```

ホストOS側の30080番ポートに httpd が接続されているので、ブラウザで localhost:30080 を確認するとApacheの ```It Works!!``` が表示される。

