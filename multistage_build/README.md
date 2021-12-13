# これは何？

Dockerfileにてマルチステージビルドを行うサンプルです。

サンプルとしてGoのアプリケーションのコンテナをマルチステージあり・なしで作ってみます。

# 実行

```sh
$ make build
$ docker image list
$ make run
$ make clean
```

# 参考情報

- [How to Minimize Go Apps Container Image](https://clavinjune.dev/en/blogs/how-to-minimize-go-apps-container-image/)
- [GoogleContainerTools distroless](https://github.com/GoogleContainerTools/distroless)
- [Dockerfileのベストプラクティス Top 20](https://sysdig.jp/blog/dockerfile-best-practices/)
- [distroless imageを実用する](https://blog.unasuke.com/2021/practical-distroless/)
- [軽量Dockerイメージに安易にAlpineを使うのはやめたほうがいいという話](https://blog.inductor.me/entry/alpine-not-recommended)
