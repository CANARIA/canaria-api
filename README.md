# Canaria-api

[![CircleCI](https://circleci.com/gh/CANARIA/canaria-api.svg?style=svg)](https://circleci.com/gh/CANARIA/canaria-api)

golang製APIサーバです。

GAE/goを使用しています。

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
export GOAPP=$HOME/google-cloud-sdk/platform/google_appengine
export PATH=$PATH:$GOPATH/bin:$GOAPP
EOF
```

### Google Cloud SDKのインストール
すべてEnterで進めてOK（bash以外を使ってる人は適宜変更）
```sh
$ curl https://sdk.cloud.google.com | bash
```

gcpコンポーネントのアップデート
```sh
$ gcloud components update
```

AppEngineSDKのインストール
```sh
$ gcloud components install app-engine-go
```

AppEngineSDKがインストール済みになっていることを確認
```sh
$ gcloud components list
```

AppEngineSDKである`goapp`コマンドに実行権限を付けておく
```sh
$ chmod +x ~/google-cloud-sdk/platform/google_appengine/goapp
```

リポジトリのクローン
```sh
$ mkdir -p $GOPATH/src/github.com/CANARIA
$ git clone git@github.com:CANARIA/canaria-api.git

# もしくは(ｇｈｑが入っていれば)
$ ghq get git@github.com:CANARIA/canaria-api.git
```

## APIサーバの起動

```sh
# (※初回のみ)依存ライブラリの管理にglideを使ってるので初回のみ先にインストール
$ brew install glide

# 依存ライブラリのDL
$ make deps

# APIサーバの起動
$ make run dev
```

`localhost:8000`でAppEngineSDKにアクセス、
`localhost:8080`でAPIサーバーにアクセスできます

### GAE

デプロイ例(直接ローカルからデプロイすることはないがメモ程度に)
```
$ gcloud app deploy --version blue app/stg.yaml --quiet --project canaria-io
# ↓はうまくいく
$ goapp deploy -application canaria-io -version blue ./app/stg.yaml
```

ターミナルから直接ログを見る
```
$ gcloud app logs tail -s stg-api --project canaria-io
```
