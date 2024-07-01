# mymodule

`mymodule` は、Go言語で書かれたサンプルプロジェクトです。

## 機能

- 整数の加算を行う `mathutils.Add` 関数
- 整数の減算を行う `mathutils.Subtract` 関数

## フォルダ構造

- cmdフォルダー： アプリケーションのエントリーポイント（main.go）を格納。
- mathutilsフォルダー： 外部に公開するパッケージを格納。
- go.modファイル： モジュールの定義とその依存関係を管理。

## インストール

このリポジトリをクローンします。

git clone https://github.com/hiroto60/mymodule.git

プロジェクトのルートディレクトリに移動します。

cd mymodule

依存関係をダウンロードします。

go mod download

## 使用方法

アプリケーションをビルドします。

go build

ビルドされたアプリケーションを実行します。

Windowsの場合：

.\mymodule.exe

macOSやLinuxの場合：

./mymodule

## ライセンス

このプロジェクトはMIT Licenseの下で公開されています。