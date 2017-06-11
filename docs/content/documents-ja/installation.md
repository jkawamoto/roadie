---
title: インストールと初期設定
---
### インストール
`roadie` のコンパイル済みバイナリは，GitHub の
[リリースページ](https://github.com/jkawamoto/roadie/releases)
から取得できます．
アーカイブを解凍して得られた実行ファイルをパスの通ったところへ置いてください．
一時的に試してみる場合は，ソースコードのあるディレクトリに置くだけでも良いです．

また，mac の場合は [Homebrew](http://brew.sh/) でもインストールできます．

```shell
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

```shell
$ roadie init
```

初期設定では，上記のプロジェクト ID の他，データの保存先やプログラムを実行するリージョンや
仮想マシンのデフォルトなどを設定します．

設定ファイルは，`roadie init` を実行したディレクトリの `roadie.yml` です．
テキスト形式なので後から編集できます．
また，`roadie config` コマンドで対話的に変更することもできます．
詳しくは [設定について](documents-ja/configuration) のページをご覧ください．
