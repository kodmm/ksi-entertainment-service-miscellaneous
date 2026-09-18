# ksi-entertainment-service-miscellaneous

「ksi-entertainment」（音楽ライブの参戦記録を残してシェアし、趣味の合う人とマッチングするサービス）の
バックエンドを構成する gRPC マイクロサービス。

## 役割

- 参戦記録（LiveRecord）のCRUD
- カード生成関連の補助処理（将来的に追加予定）

`ksi-entertainment-proto` リポジトリの `session_service.proto` などから生成される
SessionService / LiveRecordService の実装は今後のタスクで結線する。現時点では
gRPC Health Checking Protocol のみを提供する。

## 起動方法

```bash
go run .
```

デフォルトでは `50051` ポートで起動する。環境変数 `PORT` を指定すると、そのポートで起動する。

```bash
PORT=8080 go run .
```

## 構成（Clean Architecture + CQRS）

Clean Architectureの層分割に加えて、書き込み（コマンド）と読み取り（クエリ）を
分離するCQRS（Command Query Responsibility Segregation）を採用している
（Event Sourcingはmatching-serviceのみの採用で、このサービスでは使わない）。

```
internal/
├── domain/              # LiveRecord エンティティ。他レイヤーに依存しない
├── application/
│   ├── command/         # 書き込み側ハンドラー（例: CreateLiveRecordHandler）
│   └── query/           # 読み取り側ハンドラー（例: ListLiveRecordsHandler）
├── infrastructure/       # 永続化の実装。現在はインメモリ実装のみ
└── interfaces/
    └── grpc/            # gRPCサーバーの構築・サービス登録（Health Checkなど）
```

`main.go` はコンポジションルートとして、上記の実装を組み立てて結線し、
`Listen` するだけの役割に絞っている。

