---
title: ステータスとログ
description: >-
  本ドキュメントでは、クラウド環境にて実行中プログラムの状態を取得する方法並びに、ログデータの
  取得方法について説明します。
date: 2016-08-14
lastmod: 2017-10-09
slug: logging
---
実行したプログラムが問題なく動作しているか確認するためには、
ステータスのチェックやログの確認が便利です。

Roadie は次の二つのコマンドを提供しています。

- **status** は仮想マシンステータスの確認と停止を行います。
- **log** はログの表示を行います。

### status コマンド
`roadie status` コマンドは、各プログラムの実行状況を表示します。
もし、プログラムが終了済みで実行結果も削除されている場合、
`status` コマンドでは非表示になります。
計算結果の削除済みインスタンスに関する情報も取得するには `--all` フラッグを付けてください。

なお、Google Cloud Platform では、ステータス情報は一定期間が過ぎたのち自動的に削除されます。
この場合、`--all` フラッグを使用しても表示されなくなります。

`status` コマンドの別の機能は、インスタンスを途中で削除することです。
`INSTANCE` というインスタンスを何らかの理由で途中削除する場合、

```shell
$ roadie status kill INSTANCE
```

を実行してください。

なお、実行途中にインスタンスを削除した場合、
計算結果はクラウドストレージに保存されないことがあります。


### log コマンド
ログコマンドは、各インスタンスのログ出力を取得します。
`INSTANCE` というインスタンスのログを取得するには、

```shell
$ roadie log INSTANCE
```

を実行してください。

ログメッセージには、事前準備、事後プロセスの進捗に加えて、
各コマンドが標準エラー出力 `stderr` に書き出したメッセージも含まれます。

`roadie` の実行モデルでは、標準出力 `stdout` と標準エラー出力 `stderr` は
別々の意味を持ちます。
標準出力に書き出されたメッセージは、
プログラムの実行結果としてクラウドストレージに保存されます。
これらのメッセージはプログラムの実行中に外部から取得することはできませんが、
永続的に管理されます。
標準出力へ書き出されたメッセージを取得するには `roadie result` コマンドを使用します。

一方、標準エラー出力に書き出されたメッセージは、永続的なクラウドストレージに保存されるとは限りません。
その代わり、`log` コマンドを用いてインスタンスの実行中に外部から確認することができます。
標準エラー出力への書き込みは、その都度通信を発生させるため、
巨大なデータの書き込みは避けた方が良いでしょう。


#### オプション
ログコマンドには幾つかのオプションがあります。

`--no-timestamp` フラッグが付けられた場合、タイムスタンプが省略されます。

`--follow` フラッグが付けられた場合、`roadie` コマンドは新たなログの到着を待ち続けます。
このフラッグを使用した場合、`roadie` コマンドは自動では終了しなくなります。
終了させる場合には `control + c` を使用してください。

なお、クラウドプロバイダによってはログ出力が課金対象であったり、一定時間内に取得できる
ファイルサイズに制限がある場合があります。
