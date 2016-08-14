---
title: ファイルの管理
---
`roadie` は 3種類のデータをクラウドストレージ上に保存します．
それらのデータを管理するために，`source`, `data`, `result` の
3コマンドが用意されています．


### source コマンド
`source` コマンドは `roadie run` に `--local` フラッグを付けてアップロードした
ソースコードの管理を行います．

これまでにアップロードされたソースコードアーカイブの一覧を取得するには，

```sh
$ roadie source list
```

を実行します．また，その中から `FILENAME` というファイルを削除する場合，

```sh
$ roadie source delete FILENAME
```

を実行します．


### data コマンド
`data` コマンドは，スクリプトファイルの `data` セクションで利用する
データファイルの管理を行います．

あるファイル `FILENAME` を `data` セクションで利用するためにアップロードする場合，

```sh
$ roadie data put FILENAME
```

を実行してください．アップロードされたファイルの URL が表示されます．
また，これまでにアップロードしたファイルと，その URL を取得するには，

```sh
$ roadie data list --url
```

を実行してください．

アップロード済みのデータファイル `FILENAME` を削除するには，

```sh
$ roadie data delete FILENAME
```

を実行してください．


### result コマンド
`result` コマンドは計算結果としてアップロードされたファイルを管理し，
ローカル PC へのダウンロード機能を提供します．

実行が終了したインスタンスの一覧を取得するには，

```sh
$ roadie result list
```

を実行してください．また，各インスタンス `INSTANCE` がアップロードした
計算結果のファイルを調べるには，

```sh
$ roadie result INSTANCE
```

を実行してください．

`roadie result get` コマンドは，
それらのファイルのうち与えられた Glob パタンにマッチするファイルをダウンロードします．
ワイルドカードを用いてすべての結果を取得する場合，

```sh
$ roadie result get INSTANCE "*" -o ./res
```

を実行してください．`-o` フラッグによりダウンロードしたファイルの保存先を指定できます．
