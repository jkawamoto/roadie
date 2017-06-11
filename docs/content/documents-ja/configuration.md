---
title: 設定
---
設定の更新には，`roadie config` コマンドが使用できます．
このコマンドは，次のサブコマンドを提供します．

- **project** 現在の Google Cloud Platform プロジェクト ID を表示・変更します,
- **region** 現在のゾーン設定を表示・変更します,
- **machine** 現在の仮想マシンタイプを表示・変更します,

すべての設定は，カレントディレクトリの `roadie.yml` ファイルに保存されます．
このファイルはテキストファイルなので，手作業で編集できます．

### project
プロジェクト ID は Google Cloud Platform に登録されている ID です．
利用可能なプロジェクト ID の一覧は [ここ](https://console.cloud.google.com/project)
から取得できます．

現在 `roadie` に設定されているプロジェクト ID を調べるには，

```shell
$ roadie config project
```

を実行します．また，プロジェクト ID `PROJECT` を設定するには，

```shell
$ roadie config project set PROJECT
```

を実行します．


### region
Google Cloud Platform では，仮想マシンを実行するデータセンターが幾つか提供されています．
仮想マシンの利用料金はゾーンごとに異なっています．
利用可能なゾーンを取得するためには，

```shell
$ roadie config region list
```

を実行してください．
現在設定されているゾーンを調べるためには，

```shell
$ roadie config region
```

を実行してください．
なお，デフォルトは `us-central1-b` です．

最後に，ゾーンをある値 `ZONE` に変更するためには，

```shell
$ roadie config region set ZONE
```

を実行してください．

### machine
プログラムを実行する仮想マシンのスペックにはいくつかの種類が用意されています．
各マシンタイプには異なる仮想 CPU 数やメモリサイズが設定されています．
利用可能な仮想マシンタイプを取得するには，

```shell
$ roadie config machine list
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

```shell
$ roadie config machine
```

を実行してください．
また，仮想マシンタイプを `TYPE` へ変更する場合，

```shell
$ roadie config machine set TYPE
```

を実行してください．
