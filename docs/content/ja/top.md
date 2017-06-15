---
title: Roadie
description: >-
  Roadie は，Google Cloud Platform と Microsoft Azure であなたのプログラムを実行する手助けを行います．
  Roadie を用いることで，クラウドプラットフォームの詳しい仕組みを理解するために時間を割くことなく，
  簡単にプログラムを実行し計算結果を取得することができます．
date: 2016-08-14
lastmod: 2017-05-04
slug: top
---
`Roadie` は，クラウドプラットフォームであなたのプログラムを実行する手助けを行います．

`Roadie` を用いることで，クラウドプラットフォームの詳しい仕組みを理解するために時間を割くことなく，
簡単にプログラムを実行し計算結果を取得することができます．

`Roadie` は，次のクラウドプラットフォームをサポートしています．

- <i class="fa fa-google" aria-hidden="true"></i> [Google Cloud Platform](https://cloud.google.com/)
- <i class="fa fa-windows" aria-hidden="true"></i> [Microsoft Azure](https://azure.microsoft.com/)

`Roadie` では，[YAML](https://ja.wikipedia.org/wiki/YAML) 形式のスクリプトファイルを使って，
クラウドでプログラムを実行する手順を指定します．

例えば，クラウドプラットフォームを使って台湾中央研究院で公開されている[データファイル](http://mmnet.iis.sinica.edu.tw/dl/wowah/)を
カレントディレクトリにある `analyze`というプログラム<sup>[1](#program-requirement)</sup> で解析する場合，
次のようなスクリプトファイルを用意し `script.yml` という名前で保存します．

```yaml
apt:
- unrar
data:
- http://mmnet.iis.sinica.edu.tw/dl/wowah/wowah.rar
run:
- unrar x -r wowah.rar
- ./analyze WoWAH
upload:
- *.png
```

そして，次のコマンドを実行することで `Roadie` は必要なクラウド環境をセットアップしプログラムを実行します．

```shell
$ roadie run --local . script.yml
```

`Roadie` は解析プログラムが終了すると，計算結果を保存したのち不要となったクラウド環境を削除し **クラウド使用料金を節約します**．
プログラムが終了するまでコンピュータの前で待機し手動で削除処理を行う必要はありません．


### コンテンツリスト
- [インストール方法](ja/installation)
- [設定変更](ja/configuration)
- [プログラムの実行](ja/execution)
- [ステータスとログ](ja/logging)
- [ファイルの管理](ja/data)
- [タスクキュー](ja/queue)

<a name="program-requirement">1</a>: `Roadie` ではクラウドプラットフォーム上に
[Ubuntu](https://www.ubuntu.com/) 環境をセットアップしてプログラムを実行します．
そのため，実行するプログラムは 64bit Linux で実行可能なバイナリかスクリプトである必要があります．
詳細は，[プログラムの実行](ja/execution) をご覧ください．
