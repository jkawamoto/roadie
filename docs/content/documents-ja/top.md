---
title: ドキュメント
---
`roadie` は [Google Cloud Platform](https://cloud.google.com/)
でプログラムを実行する手助けをします．
具体的には，クラウドへソースコードをアップロードし，仮想マシンの作成・削除，
そしてプログラムの出力結果の管理を行います．

`roadie` を使うことで，Google Cloud Platform の詳しい仕組みを知らなくても，
また仮想マシンにログインして作業することなく簡単にクラウド環境でプログラムを実行することができます．

### 項目
- [インストール](documents-ja/installation)
- [設定変更](documents-ja/configuration)
- [プログラムの実行](documents-ja/execution)
- [ステータスとログ](documents-ja/logging)
- [ファイルの管理](documents-ja/data)

### 簡単な使用例
`roadie` では，[YAML](https://ja.wikipedia.org/wiki/YAML) 形式のスクリプトファイル
を使って，クラウドで実行する手順を指定します．
このスクリプトファイルが `script.yml` だとすると，
次のようにしてプログラムをクラウドで実行させることができます．

```sh
$ roadie run --local . --name analyze-wowah script.yml
```

このコマンドでは，`--local .` によって現在のディレクトリにあるソースコードを
クラウドにアップロードするよう指定しています．
また，後ほど参照するために `analyze-wowah` という名前を付けています．

ここで実行した `script.yml` は，次のようなフォーマットになっています．

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

このスクリプトファイルでは，次のことを指定しています．

1. プログラムの実行に必要なパッケージとして `unrar` パッケージをインストールする．
2. http://mmnet.iis.sinica.edu.tw/dl/wowah/wowah.rar
  から解析対象のデータファイルをダウンロードする．
3. `unrar` を使って取得した RAR ファイルを解凍する．
4. アップロードしたソースコードに含まれている `analyze` コマンドを使って解析を行う．
5. `analyze` コマンドが出力した `*.png` ファイルをクラウドへアップロードする．

以上の手順が無事終了すると，`roadie` は作成した仮想マシンを削除します．
そのため，仮想マシンの停止を忘れて課金が継続されることはありません．

作成した仮想マシンの状態は，次のコマンドで確認することができます．

```sh
$ roadie status
```

プログラムが終了し，アップロードされた出力結果を取得するには，
次のコマンドを利用します．

```sh
$ roadie result get analyze-wowah "*" -o ./res
```

このコマンドは，先ほど名前を付けた `analyze-wowah` の出力のうち，
"\*" にマッチする，つまりすべてのファイルを `./res` ディレクトリにコピーします．
