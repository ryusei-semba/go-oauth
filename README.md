# go-oauth

OAuth2.0認証サーバーの実装プロジェクトです。RFC 6749に準拠したトークン管理システムを提供します。

## 開発プロセス

このOAuth2.0認証サーバーは、Anthropic社のClaude 3.5 Sonnetとの対話を通じて開発されました。

### AI支援開発
- Cursor IDE上でのペアプログラミング
  - Claude 3.5 Sonnetによるコード生成
  - リアルタイムのコードレビューとリファクタリング提案
  - アーキテクチャ設計のアドバイス
- インタラクティブな開発プロセス
  - 要件定義から実装までの対話的な進行
  - エラー解決とデバッグのサポート
  - ベストプラクティスの提案と適用

## 技術スタック

### 開発環境
- Visual Studio Code
  - Go拡張機能
  - REST Client拡張機能
  - Git Graph拡張機能
- Cursor
  - AIアシスト機能
  - コード補完
  - リアルタイムエラー検出

### バックエンド
- Go 1.21
- gin-gonic/gin
- golang/oauth2
- joho/godotenv
- DuckDB (開発用データベース)

### API仕様
- OAuth 2.0
  - RFC 6749準拠
  - Bearer Token (RFC 6750)
  - JWT (RFC 7519)
- APIバージョン: v1
  - エンドポイントプレフィックス: `/api/v1`
  - Content-Type: `application/json`
  - 認証ヘッダー: `Authorization: Bearer <token>`

## アーキテクチャ

このプロジェクトはクリーンアーキテクチャを採用しています。

### ディレクトリ構成
```
.
├── .air.toml
├── .env
├── .gitignore
├── README.md
├── go.mod
├── go.sum
├── interfaces
│   ├── handler.go
│   └── route.go
├── main.go
└── usecase
    └── token_usecase.go
```

### 各ディレクトリ・ファイルの説明

#### メインディレクトリ
- `main.go`: アプリケーションのエントリーポイント。サーバーの起動と初期設定を行います。
- `.air.toml`: ホットリロード用の設定ファイル。開発時のコード変更を自動で反映します。
- `.env`: 環境変数の設定ファイル。APIキーなどの機密情報を管理します。
- `go.mod`, `go.sum`: Goの依存関係管理ファイル。

#### interfaces/
HTTPリクエストの受け付けとルーティングを担当するレイヤー
- `handler.go`: HTTPリクエストのハンドリング処理を実装
- `route.go`: エンドポイントのルーティング設定を管理

#### usecase/
アプリケーションのビジネスロジックを実装するレイヤー
- `token_usecase.go`: OAuth認証に関するトークンの処理を実装

### 今後の改善計画

#### 追加予定のディレクトリ/レイヤー
1. `domain/`: ビジネスエンティティとドメインルールを定義
   - エンティティの構造体定義
   - ドメインのインターフェース定義

2. `repository/`: データの永続化層を実装
   - データベースアクセス
   - 外部APIとの通信処理

3. `infrastructure/`: 技術的な実装の詳細
   - データベース接続
   - 外部サービスクライアント
   - ミドルウェア

#### リファクタリングポイント
1. 依存関係の方向
   - 外側のレイヤー（interfaces）から内側のレイヤー（domain）への依存のみとする
   - インターフェースを活用した依存性の逆転

2. ユースケースの分離
   - 現在の`token_usecase.go`をより細かい責務に分割
   - インターフェースとその実装を明確に分離

3. エラーハンドリング
   - ドメイン固有のエラー型の定義
   - レイヤー間のエラー変換処理の実装


# cursor-fe
