---
title: インストールと初期設定
description: >-
  Roadie のインストール方法と初期設定の方法について説明します．
  Roadie は Google Cloud Platform と Microsoft Azure に対応しています．
  それぞれのクラウドプラットフォームで必要となる設定は異なります．
  この文書では，プラットフォーム別にスクリーンショット付きで必要な情報の収集方法を説明します．
date: 2016-08-14
lastmod: 2017-10-06
slug: installation
---
### インストール
`roadie` のコンパイル済みバイナリは，GitHub の
[リリースページ](https://github.com/jkawamoto/roadie/releases)
から取得できます．
アーカイブを解凍して得られた実行ファイルをパスの通ったところへ置いてください．
一時的に試してみる場合は，ソースコードのあるディレクトリに置くだけでも良いです．

[Homebrew](http://brew.sh/) または [Linuxbrew](http://linuxbrew.sh/) を
お使いの場合は，下記のコマンドでもインストールできます．

```shell
$ brew tap jkawamoto/roadie
$ brew install roadie
```


### 初期設定
`Roadie` は，現在のところ [Google Cloud Platform](https://cloud.google.com/) と
[Microsoft Azure](https://azure.microsoft.com/) をサポートしています．
次のコマンドを実行することで，どちらのクラウドプラットフォームを利用するのかを指定し，必要な情報を登録します．

```shell
$ roadie init
```

#### Google Cloud Platform
Google Cloud Platform では，プロジェクトという単位でリソースを管理しています．
どのプロジェクトを利用してプログラムを実行するのか `Roadie` に設定する必要があります．

利用可能なプロジェクト ID は[クラウドコンソール](https://console.cloud.google.com/project)から調べることができます．
なお，**プロジェクト名** ではなく **プロジェクト ID** を使うことに注意してください．

<img src="img/gcp-projects.png" style="width: 80%;" title="Google Cloud Platform プロジェクト一覧"/>

`roadie init` を実行し，クラウドプラットフォームとして Google Cloud Platform (gcp) を選択すると，
プロジェクト ID を尋ねられるので入力してください．
プロジェクト ID を入力すると認証プロセスが始まります．
`Roadie` を実行している端末にて Web ブラウザで表示された URL を開いてください．
Google へログインし，`Roadie` へのアクセス許可を与えてください．

認証に成功すると設定ファイル `roadie.yml` が作成されます．
その他の設定については[設定について](documents-ja/configuration) のページをご覧ください．

#### Microsoft Azure
Microsoft Azure では，ユーザアカウントは一つの Active Directory に所属しています．
また，一つの Active Directory には複数のサブスクリプションつまり課金体系が登録され得ます．
したがって，初期設定では `Roadie` を利用するユーザの Active Directory の ID (テナント ID) と
サブスクリプション ID の二つの情報を指定します．

テナント ID の取得には，まずポータルサイトから Azure Active Directory を選びます．

<img src="img/active-directory.png" style="width: 80%;" title="Azure Active Directory"/>

表示される Active Direcotry に関する情報の中からプロパティを選ぶと，
ディレクトリ ID という名称でテナント ID が表示されています．

<img src="img/tenant-id.png" style="width: 80%;" title="テナント ID"/>

サブスクリプション ID の取得は，同じくポータルサイトから課金を選びます．

<img src="img/subscription.png" style="width: 80%;" title="課金"/>

表示される課金情報の中にサブスクリプション ID の一覧があります．

<img src="img/subscription2.png" style="width: 80%;" title="サブスクリプション ID"/>

`roadie init` を実行し，クラウドプラットフォームとして Microsoft Azure (azure) を選択すると，
テナント ID とサブスクリプション ID を尋ねられるので入力してください．
入力し終わると認証プロセスが始まります．
Web ブラウザで表示された URL を開き同じく表示されたコードを入力してください．
Microsoft へログインし，`Roadie` へのアクセス許可を与えてください．

認証に成功すると設定ファイル `roadie.yml` が作成されます．
その他の設定については[設定について](documents-ja/configuration) のページをご覧ください．
