# BusyBox とは

俗に組み込み系の十徳ナイフと呼ばれる超軽量Linux。

フロッピー１枚分程度の容量で、基本的なコマンドが２００個くらい入っている。

なにかの調整作業をしたいときなどに重宝する。

# イメージのサイズ

```sh
$ docker image list | grep busybox
busybox                                                    latest    7138284460ff   10 days ago     1.24MB
```

# 参考情報

- [組み込みLinuxで際立つ「BusyBox」の魅力](https://monoist.itmedia.co.jp/mn/articles/0802/04/news114.html)
- [組み込みLinuxの十徳ナイフ。多数のコマンドを集めたBusyBoxの概要と使い方](https://kakurasan.blogspot.com/2015/08/busybox-features-usage.html)
