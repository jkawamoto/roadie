---
title: インストールと初期設定
---
### インストールの前に
`roadie` は [Google Cloud Platform](https://cloud.google.com/)
でプログラムを実行します．
そのため，Google Cloud Platform が利用可能になっている必要があります．

また，`roadie` は Google の [Cloud SDK](https://cloud.google.com/sdk/) を利用します．
[ここ]((https://cloud.google.com/sdk/)) から取得できるので，
インストールし初期設定を行ってください．
`gcloud auth list` を実行してみて，認証済みのアカウント名（メールアドレス）が表示されれば
初期化は完了しています．


### インストール
`roadie` のコンパイル済みバイナリは，GitHub の
[リリースページ](https://github.com/jkawamoto/roadie/releases)
から取得できます．
アーカイブを解凍して得られた実行ファイルをパスの通ったところへ置いてください．
一時的に試してみる場合は，ソースコードのあるディレクトリに置くだけでも良いです．

また，Go 言語の環境設定が済んでいる場合，`go get` コマンドでもインストールできます．

```sh
$ go get github.com/jkawamoto/roadie
```

この場合，実行ファイルは `$GOPATH/bin` に保存されます．

また，mac の場合は [Homebrew](http://brew.sh/) でもインストールできます．

```sh
$ brew tap jkawamoto/roadie
$ brew install roadie
```


### 初期設定
Google Cloud Platform では，プロジェクトという単位でリソースを管理しています．
`roadie` を利用する前に，どのプロジェクトを利用するか `roadie` に設定する必要があります．
なお，利用可能なプロジェクト ID は [ここ](https://console.cloud.google.com/project)
から調べることができます．
**プロジェクト名** ではなく **プロジェクト ID** を使うことに注意してください．

`roadie` の初期設定には，次のコマンドを利用します．

```sh
$ roadie init
```

初期設定では，上記のプロジェクト ID の他，データの保存先やプログラムを実行するリージョンや
仮想マシンのデフォルトなどを設定します．

設定ファイルは，`roadie init` を実行したディレクトリの `.roadie` です．
テキスト形式なので後から編集できます．
また，`roadie config` コマンドで対話的に変更することもできます．
詳しくは [設定について](documents-ja/configuration) のページをご覧ください．
