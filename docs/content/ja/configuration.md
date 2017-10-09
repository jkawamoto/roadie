---
title: 設定について
description: >-
  本ドキュメントでは、Google Cloud Platform と Microsoft Azure それぞれをバックエンド
  として使用する場合における Roadie の設定について説明します。
date: 2016-08-14
lastmod: 2017-10-10
slug: configuration
---
Roadie の設定は [YAML](https://ja.wikipedia.org/wiki/YAML) 形式のテキストファイルなので、
テキストエディタ等で直接編集可能です。
加えて、一般的な設定項目を変更する `roadie config` コマンドも提供しています。

本セクションでは、はじめに `roadie config` コマンドによる設定変更方法について紹介したのち、
直接 YAML ファイルを編集する方法について紹介します。

### roadie config コマンド
設定の更新には、`roadie config` コマンドが使用できます。
このコマンドは、次のサブコマンドを提供します。

- project: プロジェクト ID を表示・変更します,
- region: プログラムを実行するリージョン設定を表示・変更します,
- machine: プログラムを実行する仮想マシンの種類を表示・変更します,

デフォルトの設定ファイルはカレントディレクトリの `roadie.yml` です。
他の設定ファイルを使用する場合、roadie とコマンド名の間に `-c` オプションを使って指定できます。
例えば、

```shell
$ roadie -c another.yml config project
```

とした場合、`another.yml` を設定ファイルとして使用します。


#### project
現在設定されているプロジェクト ID を調べるには、

```shell
$ roadie config project
```

を実行します。また、`PROJECT` を新しいプロジェクト ID として設定するには、

```shell
$ roadie config project set PROJECT
```

を実行します。


#### region
一般的に、クラウド利用料金はリージョンごとに異なっています。
利用可能なリージョンを取得するためには、

```shell
$ roadie config region list
```

を実行してください。
現在設定されているリージョンを調べるためには、

```shell
$ roadie config region
```

を実行してください。

最後に、リージョンをある値 `REGION` に変更するためには、

```shell
$ roadie config region set REGION
```

を実行してください。

#### machine
プログラムを実行する仮想マシンのスペックにはいくつかの種類が用意されています。
各マシンタイプには異なる仮想 CPU 数やメモリサイズが設定されています。
利用可能な仮想マシンタイプを取得するには、

```shell
$ roadie config machine list
```

を実行してください。

利用可能な仮想マシンタイプはリージョンによって異なる可能性があります。
仮想マシンタイプを設定する前に、リージョンの設定を行ってください。

現在設定されている仮想マシンタイプを取得するには、

```shell
$ roadie config machine
```

を実行してください。
また、仮想マシンタイプを `TYPE` へ変更する場合、

```shell
$ roadie config machine set TYPE
```

を実行してください。

### 設定ファイルの直接編集
先に述べたの通り、Roadie の設定は YAML 形式のテキストファイルです。
この節では、この設定ファイルがどのような要素を持つのかについて説明します。

なお、Roadie 設定ファイルのルート要素は、どのクラウドプロバイダを利用しているのかによって異なります。

#### Google Cloud Platform
[Google Cloud Platform](https://cloud.google.com/) を利用している場合、
ルート要素は `gcp` になり、以下の要素を持ちます。

```yaml
gcp:
  project: <project ID>
  bucket: <bucket ID>
  zone: us-central1-b
  machine_type: n1-standard-1
  disk_size: 10
```

- `project`: 初期設定の節で述べたプロジェクト ID を設定します。
- `bucket`: 実行結果を保存する [Cloud Storage](https://cloud.google.com/storage/)
  のバケット ID を設定します。
  このバケット ID も全世界でユニークである必要があります。
  デフォルトでは、プロジェクト ID と同じ値を使用します。
  また、バケットが作成されていない場合は必要に応じて作成されます。
- `zone`: プログラムを実行する仮想マシンを作成するリージョン（ゾーン）を指定します。
  デフォルトでは、時間あたりの費用が最も安いリージョンの一つである us-central1-b が選ばれています。
- `machine_type`: プログラムを実行する仮想マシンの種類を指定します。
  デフォルトでは、時間あたりの費用が安い n1-standard-1 が選ばれています。
- `disk_size`: プログラムを実行する仮想マシンのディスクサイズを GB 単位で指定します。
  デフォルト値は 10 (GB) です。

#### Microsoft Azure
[Microsoft Azure](https://azure.microsoft.com/) を利用している場合、
ルート要素は `azure` になり、以下の要素を持ちます。

```yaml
azure:
  tenant_id: <ディレクトリ ID (テナント ID)>
  subscription_id: <サブスクリプション ID>
  project_id: <プロジェクト ID>
  location: westus2
  machine_type: Standard_A2
  os:
    publisher_name: Canonical
    offer: UbuntuServer
    skus: "17.04"
    version: latest
```

- `tenant_id`: 初期設定の節で述べたディレクトリ ID を設定します。
- `subscription_id`: 初期設定の節で述べたサブスクリプション ID を設定します。
- `project_id`: 初期設定の節で述べたプロジェクト ID を設定します。
  Roadie の実行に必要なリソースはここで指定した名称でグループ化されます。
- `location`: プログラムの実行する仮想マシンを作成するリージョンを指定します。
  デフォルトは westus2 が選ばれています。
- `machine_type`: プログラムを実行する仮想マシンの種類を指定します。
  デフォルトでは、時間あたりの費用が安い Standard_A2 が選ばれています。
- `os`: プログラムの実行に利用する OS 情報を設定します。
  現在のところデフォルトで設定されている値以外はサポートしていません。
