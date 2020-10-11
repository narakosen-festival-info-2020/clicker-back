# clicker-back
Back-end of Clicker game

## API

### client -> server

- 前回送信時からの差分クリック回数を送信
    - 正直クリック毎に1ずつ送信しても問題なさげ

```json
{
  "count": 1
}
```

### server -> client

- 現在の総クリック数を送信

```json
{
  "count": 300
}
```

## For Developers

### 環境構築

1. `> env.sh` or `> env.bat`
2. `> mkdir -p app/tmp/pids`
3. `> up`
4. `> run`以降に実行したいコマンドを入力
    - `go`コマンドはそのままで実行可能

### Docker関連

- `docker/Dockerfile`
    - 本番環境用(未完成)
- `docker/local.Dockerfile`
    - 開発環境用
- docker-compose
    - develop
        - 諸々のコマンド実行用
        - `run`はここで実行
    - app
        - `main.go`が動作してるやつ
        - 環境自体はdevelopと同じ
    - nginx
        - リバースプロキシ
        - 使わないかもしれない
    - postgres
        - DB
        - 使わないかもしれない
