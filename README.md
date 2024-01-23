# Introduction
Go学習用環境

# 環境構築
Macを想定

```shell
$ make install
```

# コマンド実行
## main.goを実行
```shell
$ make main

# or...

$ docker compose exec go bash -c 'go run /go/src/app/main.go'
```

## コンテナにアタッチ
```shell
$ make go

# or...

$ docker compose exec go bash
```