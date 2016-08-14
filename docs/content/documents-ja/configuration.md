---
title: 設定
---
設定の更新には，`roadie config` コマンドが使用できます．
このコマンドは，次のサブコマンドを提供します．

- **project** 現在の Google Cloud Platform プロジェクト ID を表示・変更します,
- **zone** 現在のゾーン設定を表示・変更します,
- **type** 現在の仮想マシンタイプを表示・変更します,
- **bucket** 現在の Google Cloud Storage バケット名を表示・変更します．

すべての設定は，カレントディレクトリの `.roadie` ファイルに保存されます．
このファイルはテキストファイルなので，手作業で編集できます．

### project
プロジェクト ID は Google Cloud Platform に登録されている ID です．
利用可能なプロジェクト ID の一覧は [ここ](https://console.cloud.google.com/project)
から取得できます．

現在 `roadie` に設定されているプロジェクト ID を調べるには，

```sh
$ roadie config project
```

を実行します．また，プロジェクト ID `PROJECT` を設定するには，

```sh
$ roadie config project set PROJECT
```

を実行します．

プロジェクト ID は，`roadie` の他のコマンドが Google Cloud Platform に
アクセスするために使用されます．
もし，正しいプロジェクト ID を設定しているにも関わらず，
他のコマンドが失敗する場合，正しく認証されていることを確認してください．
`gcloud auth list` は，認証状況を出力します．


### zone
Google Cloud Platform では，仮想マシンを実行するデータセンターが幾つか提供されています．
仮想マシンの利用料金はゾーンごとに異なっています．
利用可能なゾーンを取得するためには，

```sh
$ roadie config zone list
```

を実行してください．
現在設定されているゾーンを調べるためには，

```sh
$ roadie config zone
```

を実行してください．
なお，デフォルトは `us-central1-b` です．

最後に，ゾーンをある値 `ZONE` に変更するためには，

```sh
$ roadie config zone set ZONE
```

を実行してください．

### type
プログラムを実行する仮想マシンのスペックにはいくつかの種類が用意されています．
各マシンタイプには異なる仮想 CPU 数やメモリサイズが設定されています．
利用可能な仮想マシンタイプを取得するには，

```sh
$ roadie config type list
```

を実行してください．
また，仮想マシンタイプと利用料金の詳細は，
[こちら](https://cloud.google.com/compute/pricing)
を参照してください．
利用可能な仮想マシンタイプはゾーンによって異なる可能性があります．
仮想マシンタイプを設定する前に，ゾーンの設定を行ってください．

デフォルトでは，1つの仮想 CPU と 3.75 GB のメモリを使用する，
`n1-standard-1` が選ばれています．
現在設定されている仮想マシンタイプを取得するには，

```sh
$ roadie config type
```

を実行してください．
また，仮想マシンタイプを `TYPE` へ変更する場合，

```sh
$ roadie config type set TYPE
```

を実行してください．

### bucket
バケットは，ソースコードを含め `roadie` が使用するデータを保存する
クラウドストレージの場所です．
バケットは全世界でユニークなバケット名で識別されます．
任意のバケット名を使用可能ですが，プロジェクト ID と同じものを設定しておく方が良いでしょう．
なお，デフォルトではプロジェクト ID と同じ値が設定されます．

現在設定されているバケット名を調べるには，

```sh
$ roadie config bucket
```

を実行します．また，バケット名として `NAME` を使用する場合，

```sh
$ roadie config bucket set NAME
```

を実行します．
