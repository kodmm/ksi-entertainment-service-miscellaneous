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
