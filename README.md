# clicker-back
Back-end of Clicker game

## For Developers

### 環境構築

1. `> env.sh` or `> env.bat`
2. `> up`
3. `> run`以降に実行したいコマンドを入力
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
    - postgres
        - DB
        - 使わないかもしれない
