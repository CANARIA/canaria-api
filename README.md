# Canaria-api

APIサーバです。

*事前にdocker for macをインストールしておくこと*

## dockerコンテナの立ち上げ

```
# コンテナの立ち上げ
$ docker-compose up -d

# コンテナの確認
$ docker-compose ps
```

立ち上がるコンテナは以下の通り
- APIコンテナ (:5000)
- MySQLコンテナ (:3306)

mysqlには`root`/`password`でログインできる<br>
データベースは`canaria`

## DBのマイグレーション
```
$ make migrate
```

## APIサーバの起動

```
# 依存ライブラリのDL
$ make deps

# APIサーバの起動
$ go run server.go
```

`localhost:5000`でAPIサーバーにアクセスできます

