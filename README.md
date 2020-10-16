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

### Make

#### 開発用
**make dev-image**  
開発用のDockerイメージをBuild  

**make dev-run**  
Buildしたイメージを実行する。  

**make dev-stop**  
実行されているコンテナを削除。  

**make dev-logs**  
ログ取得。  

**make main-image**  
本番用のDockerイメージをBuild  

**make main-run**  
Buildしたイメージを実行する。  

**make main-stop**  
実行されているコンテナを削除。  

**make main-logs**  
ログ取得。  