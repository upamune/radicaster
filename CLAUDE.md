# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

RadicastはRadikoから録音した音声ファイルをPodcast配信するためのGo製Webアプリケーション。Cron式でスケジューリングされた録音タスクを実行し、HTTP経由でRSSフィードとして配信する。

## 開発コマンド

### ビルド
```bash
make build
```
- `dist/radicaster` にバイナリが生成される

### 開発環境での起動（ホットリロード）
```bash
make watch
```
- Airを使用してファイル変更を監視し自動リビルド
- `-debug -config ./radicaster.yaml` 引数で起動
- `.air.toml` で設定を管理

### テスト実行
```bash
go test ./...
```

### 起動方法（本番）
```bash
radicaster -baseurl http://localhost:3333 -config ./radicaster.yaml -targetdir ./output
```

重要なフラグ：
- `-baseurl`: サーバーのベースURL（RSS内のURLに使用）
- `-config`: 設定ファイルパス（YAMLまたはURL）
- `-targetdir`: 録音ファイルの保存先ディレクトリ
- `-basicauth`: Basic認証（`user:password` 形式）
- `-radikoemail`, `-radikopassword`: Radikoプレミアム認証
- `-debug`, `-trace`: ログレベル制御

### 重要な起動時の注意
READMEに記載の通り、ファイルディスクリプタ数の制限を増やす必要がある：
```bash
ulimit -n 16384
```

## アーキテクチャ

### パッケージ構成

- **cmd/radicaster/main.go**: エントリーポイント。HTTP サーバー、録音スケジューラー、ファイル監視の初期化と起動
- **config**: YAML設定ファイルの読み込み・パース・バリデーション
- **record**: Radikoからの録音機能とCronスケジューラー管理
- **podcast**: 録音ファイルからPodcast RSSフィードを生成
- **http**: Echo ベースのHTTPサーバー（RSS配信、設定UI、Basic認証）
- **radikoutil**: Radikoクライアントの初期化とラッパー
- **ffmpeg**: AACファイルの結合処理
- **metadata**: 音声ファイルへのID3タグ書き込み・読み込み
- **timeutil**: 日本時間(JST)や曜日計算のユーティリティ

### 主要な処理フロー

#### 1. 録音処理（record パッケージ）
- `Recorder` が `gocron` でCronスケジュールを管理
- 番組指定録音（`Record`）と全録（`RecordAll`）の2モード
- Radiko APIから番組情報とM3U8プレイリストを取得
- AACチャンクを並列ダウンロード → ffmpegで結合 → メタデータ書き込み
- 番組指定録音時、エリアIDを指定することでエリア外の番組も録音可能（プレミアム必須）

#### 2. Podcast配信（podcast パッケージ）
- `Podcaster` がファイルシステムを監視（`fsnotify`）
- ファイル変更を検知すると `Sync()` でRSSフィード再生成
- パスごとにエピソードをグループ化（メタデータの `Path` フィールドベース）
- `?since=30d` クエリパラメータで期間フィルタリング対応（1y, 6m, 30d, 24h形式）

#### 3. HTTP API（http パッケージ）
- `GET /rss.xml`: 全エピソードのRSS（`?since` パラメータでフィルタ可能）
- `GET /:program_path/rss.xml`: 番組別RSS（`?since` パラメータでフィルタ可能）
- `GET /zenroku/:program_path/rss.xml`: 全録番組のRSS（`?since` パラメータでフィルタ可能）
- `GET /config`: 設定表示（HTML/JSON/YAML）
- `PUT /config`: 設定更新（JSON/YAML）
- `GET /static/*`: 録音ファイル配信（Basic認証スキップ）
- Basic認証は静的ファイル配信以外の全エンドポイントに適用

### 重要な設計上の注意点

#### Radikoクライアントの初期化
`record.go:138, 267` のコメントにある通り、**Radikoクライアントは毎回初期化が必要**。再利用すると認証エラーになる。

#### 設定ファイル管理
- 設定はメモリ上で管理され、`Recorder.RefreshConfig` で動的更新可能
- `-config` でファイルパスを指定した場合、設定更新時にファイルも自動保存
- `-configurl` でURL指定も可能（読み取り専用）

#### 並行処理とロック
- `Recorder.config` と `Recorder.scheduler` は `sync.RWMutex` で保護
- `Podcaster.feedMap` も `sync.RWMutex` で保護
- AACダウンロードは `sourcegraph/conc/pool` で並列化

#### 全録モードの仕組み
- Zenroku（全録）は指定したステーションの全番組を録音
- `RecordAll` 関数がステーションごとに並列処理
- 番組ごとではなくステーション単位で並列化して負荷制御

#### メタデータとパス設計
- 音声ファイルのメタデータに `Path` と `ZenrokuMode` を保存
- `Path` でRSSフィードのグルーピングを制御
- Zenrokuモードの場合、パスに `zenroku/` プレフィックスが付く

## テストに関して

- テストファイルは各パッケージ内に `*_test.go` として配置
- 現在テストがあるパッケージ：`config`, `podcast`, `record`, `timeutil`
- 単一パッケージのテスト実行：`go test ./config -v`

## ログレベル

- デフォルト: Info
- `-debug`: Debug
- `-trace`: Trace
- zerologを使用（構造化ログ、スタックトレース対応）
