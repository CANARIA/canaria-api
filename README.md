# Canaria-api

golang製APIサーバです。

*事前にdocker for macをインストールしておくこと*

[ghq入れておくと幸せになるかも](http://suzumi.hatenablog.com/entry/2016/10/27/130338)

## 準備

Goを入れてない場合はインストールする
```sh
$ brew install go
```

GOPATHの設定（任意の場所でいいけど`dev`だと使いやすいかも）
```sh
# ホームディレクトリにdevディレクトリを作ってない場合は作る
$ mkdir ~/dev

#zsh使ってる人は`.zshrc`にGOPATHの設定
$ cat <<EOF >> ~/.bash_profile
export GOPATH=$HOME/dev
export GOROOT=/usr/local/opt/go/libexec # brewからインストールした場合
export PATH=$PATH:$GOPATH/bin
EOF
```

リポジトリのクローン
```sh
$ mkdir -p $GOPATH/src/github.com/CANARIA
$ git clone git@github.com:CANARIA/canaria-api.git

# もしくは(ｇｈｑが入っていれば)
$ ghq get git@github.com:CANARIA/canaria-api.git
```

## dockerコンテナの立ち上げ

docker for macを入れてない場合は先にインストールする

```sh
# コンテナの立ち上げ（プロジェクトルートで）
$ docker-compose up -d mysql redis redis-commander

# コンテナの確認
$ docker-compose ps
```

立ち上がるコンテナは以下の通り
- <s>APIコンテナ (:5000)</s>
- MySQLコンテナ (:3306)
- Redisコンテナ（:6379）
- Redis Commanderコンテナ（:8081）

mysqlには`root`/`password`でログインできる<br>
データベースは`canaria`
コマンドから使う場合は
```
$ mysql -h 127.0.0.1 -u root -p password
```

redisへは`redis-cli`コマンドで繋がる<br>
GUIで確認したい場合は`localhost:8081`でRedis Commanderが開ける


## DBのマイグレーション
```sh
$ make migrate
```

## APIサーバの起動

```sh
# (※初回のみ)依存ライブラリの管理にglideを使ってるので初回のみ先にインストール
$ brew install glide

# (※初回のみ)ログ出力先/var/log/canariaを作成してない場合は作る
$ sudo mkdir /var/log/canaria
$ sudo chmod -R 777 /var/log/canaria

# 依存ライブラリのDL
$ make deps

# 環境変数DOCKER_PASSWORDに設定しないとAPIサーバが立ち上がったときにDBコネクションエラーが出る
$ export DOCKER_PASSWORD=password

# APIサーバの起動
$ make run dev
```

`localhost:5000`でAPIサーバーにアクセスできます

## デバッグ

```sh
# すべてのコンテナのログが混ざって出る
$ docker-compose logs

# 特定コンテナのログのみを見たい(以下の例はmysqlコンテナのログを見る)
# ※コンテナ名ではなくサービス名を指定すること
$ docker-compose logs mysql

# オプションの説明
# f: ストリーミングでログを出力する
# t: 時間を含めた詳細なログを出力する
$ docker-compose logs -ft mysql
```

