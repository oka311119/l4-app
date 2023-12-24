# l4

## フロントエンド

bun vite react
CloudFront S3

## バックエンド(Command)

gin golang
Fargate

## バックエンド(Query)

https://github.com/zhashkevych/go-clean-architecture

## フォルダ構成

``` txt
project-root/
│
├── backend/                    # バックエンド用のGolangコード
│   ├── command/                # CQRS Command
|   |   |── cmd/                # アプリケーションのエントリポイント
|   |   ├── internal/           # 内部パッケージ
|   |   │   ├── api/            # APIハンドラーとルーティング
|   |   │   ├── model/          # データモデルとビジネスロジック
|   |   │   └── storage/        # データベースアクセスとストレージロジック
|   |   ├── pkg/                # 外部から使用可能なパッケージ（共通ユーティリティなど）
|   |   └── Dockerfile          # Fargate用のDockerファイル
│   └── query/                # CQRS Query
│
├── frontend/                   # フロントエンド用のReact + Viteコード
│   ├── public/                 # 静的ファイル（index.html, faviconなど）
│   ├── src/                    # Reactコンポーネントとロジック
│   │   ├── components/         # 再利用可能なコンポーネント
│   │   ├── pages/              # 各ページのコンポーネント
│   │   ├── app.jsx             # アプリケーションのルートコンポーネント
│   │   └── main.jsx            # エントリポイント
│   ├── vite.config.js          # Viteの設定ファイル
│   └── package.json            # npm依存関係とスクリプト
│
├── infra/                      # CloudFormationテンプレートやインフラ関連のスクリプト
│   ├── cloudformation/         # CloudFormationのテンプレートファイル
│   │   ├── network.yml         # VPCやサブネットなどのネットワークリソース
│   │   ├── compute.yml         # FargateやLambdaのリソース
│   │   ├── database.yml        # DynamoDBなどのデータベースリソース
│   │   └── frontend.yml        # S3, CloudFrontのフロントエンドリソース
│   └── scripts/                # デプロイメントやセットアップ用のスクリプト
│
├── .gitignore                  # Gitの無視ファイル設定
├── README.md                   # プロジェクトの説明書
└── Makefile                    # ビルドやデプロイを簡素化するためのMakefile
```
