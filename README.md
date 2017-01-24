# Canaria-api

golang製APIサーバです。

*事前にdocker for macをインストールしておくこと*

[ghq入れておくと幸せになるかも](http://suzumi.hatenablog.com/entry/2016/10/27/130338)

## 準備

GOPATHの設定（任意の場所でいいけど`dev`だと使いやすいかも）
```sh
#zsh使ってる人は`.zshrc`
$ cat <<EOF >> ~/.bash_profile
export GOPATH=$HOME/dev
EOF
```

フォークしてリポジトリのクローン

```sh
$ mkdir -p $GOPATH/src/github.com/CANARIA
$ git clone [フォークしたリポジトリパス]

# もしくは(ｇｈｑが入っていれば)
$ ghq get [フォークしたリポジトリパス]
```

## dockerコンテナの立ち上げ

```sh
# コンテナの立ち上げ
$ docker-compose up -d

# コンテナの確認
$ docker-compose ps
```

立ち上がるコンテナは以下の通り
- APIコンテナ (:5000)
- MySQLコンテナ (:3306)
- Redisコンテナ（:6379）
- Redis Commanderコンテナ（:8081）

mysqlには`root`/`password`でログインできる<br>
データベースは`canaria`

redisへは`redis-cli`コマンドで繋がる<br>
GUIで確認したい場合は`localhost:8081`でRedis Commanderが開ける


## DBのマイグレーション
```sh
$ make migrate
```

## APIサーバの起動

```sh
# 依存ライブラリのDL
$ make deps

# APIサーバの起動
$ go run server.go
```

`localhost:5000`でAPIサーバーにアクセスできます

