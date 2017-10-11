---
title: ファイルの管理
description: >-
  本ドキュメントでは、ユーザプログラム並びにユーザプログラムが使用するデータファイルの扱い、
  またクラウド環境にて作成された実行結果の取得方法について説明します。
date: 2016-08-14
lastmod: 2017-10-09
slug: data
---
Roadie は

- ソースコード
- データファイル
- 実行結果

という 3種類のデータをクラウドストレージ上に保存します。
それらのデータを管理するために、`source`, `data`, `result` の 3 つのコマンドが用意されています。


### ソースコード
`roadie run` コマンドに `--local` フラッグを付けて実行した場合にアップロードされたソースコードは、
`roadie://source/` 以下に保存されます。

`source` コマンドはこのソースコードの管理を行います。

#### list
これまでにアップロードされたソースコードアーカイブの一覧を取得するには、

```shell
$ roadie source list
```

を実行します。

#### delete
これまでにアップロードされたソースコードアーカイブの中から `FILENAME` というファイルを削除する場合、

```shell
$ roadie source delete FILENAME
```

を実行します。

#### get
これまでにアップロードされたソースコードアーカイブの中から `FILENAME` というファイルをダウンロードする場合、

```shell
$ roadie source get FILENAME
```

デフォルトではカレントディレクトリにダウンロードされますが、別のディレクトリ例えば `~/path` に保存する場合は、

```shell
$ roadie source get -o ~/path FILENAME
```

と `-o` フラグを使用します。


#### put
`roadie run` コマンドを使わずにソースコードをアップロードすることもできます。

例えば、`~/source` 以下にあるファイルを保存し、 `source.tar.gz` という名前で参照する場合、

```shell
$ roadie source put ~/source source.tar.gz
```

を実行します。


### データファイル
データファイルはスクリプトファイルの `data` セクションで利用するデータファイルを保存しておく場所です。
ファイルは、`roadie://data/` 以下に保存されます。
データファイルを管理するコマンドとして、`data` コマンドが用意されています。

#### put
あるファイル `FILENAME` を `data` セクションで利用するためにアップロードする場合、

```shell
$ roadie data put FILENAME
```

を実行してください。アップロードされたファイルの URL が表示されます。

#### list
これまでにアップロードしたファイルと、その URL を取得するには、

```shell
$ roadie data list --url
```

を実行してください。

#### delete
アップロード済みのデータファイル `FILENAME` を削除するには、

```shell
$ roadie data delete FILENAME
```

を実行してください。


#### get
アップロード済みのデータファイル `FILENAME` をダウンロードする場合、

```shell
$ roadie data get FILENAME
```

を実行してください。
デフォルトではカレントディレクトリにダウンロードされますが、別のディレクトリ例えば `~/path` に保存する場合は、

```shell
$ roadie data get -o ~/path FILENAME
```

と `-o` フラグを使用します。


### 実行結果
クラウド上で実行したプログラムが標準出力 `stdout` に書き出したメッセージと、
スクリプトの `upload` セクションで指定したファイルは、
`roadie://result/{インスタンス名}` 以下に保存されます。

これらのファイルにアクセスするために、`result` コマンドが用意されています。

#### list
実行が終了したインスタンスの一覧を取得するには、

```shell
$ roadie result list
```

を実行してください。また、各インスタンス `INSTANCE` がアップロードした
計算結果のファイルを調べるには、

```shell
$ roadie result list INSTANCE
```

を実行してください。

#### get
`roadie result get` コマンドは、
実行結果ファイルのうち与えられた Glob パタンにマッチするファイルをダウンロードします。
ワイルドカードを用いてすべての結果を取得する場合、

```shell
$ roadie result get INSTANCE "*" -o ./res
```

を実行してください。
デフォルトではカレントディレクトリにダウンロードされますが、
例のように `-o` フラッグによりダウンロードしたファイルの保存先を指定できます。


#### delete
`delete` コマンドは、
実行結果ファイルのうち与えられた Glob パタンにマッチするファイルを削除します。

例えば、次の例は拡張子が `png` であるファイルを削除します。

```shell
$ roadie result delete INSTANCE "*.png"
```

また、Glob パタンを省略した場合すべての実行結果ファイルが削除されます。
なお、**すべてのファイルを削除するとログも取得できなくなります**。


#### show
`show` コマンドは、標準出力に書き出されたメッセージを表示します。
次の例では、`INSTANCE` において標準出力に書き出された全てのメッセージを表示します。

```shell
$ roadie result show INSTANCE
```

特に、 スクリプトファイルの `run` セクションに書かれたコマンドのうち、
`i` 番目のコマンドが標準出力に書き出したメッセージのみを取得したい場合は、

```shell
$ roadie result show INSTANCE i
```

のように番号を指定します。
