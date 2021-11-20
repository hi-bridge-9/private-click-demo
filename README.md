# private-click-demo

1. ローカルで使用するIPアドレスを追加する

```sh:
$ sh script/ip.sh
```

2. 手順1で追加したIPアドレスを広告配信業者、メディア、広告主サイトに割り当てる

```sh:
$ sudo vi /etc/hosts
```

```sh:/etc/hosts
# 以下の情報を追記する
127.0.0.2       ad-deliver.test
127.0.0.3       publisher.test
127.0.0.4       advertiser.test
```

3. 広告配信業者、メディア、広告主サイトのコンテナを立ち上げる

```sh:
$ cd build
$ docker-compose up -d
```

4. ブラウザで「[http://publisher.test/top](http://publisher.test/top)」にアクセスして、広告のクリックなどをしてみる
