# IP-notificater

現在のグローバルIPを取得して、変更があったらSlackにPOSTします

## ビルド

### 環境
- Golang
- make

### 準備

ビルドする前にカレントディレクトリにSlack-APIのURLが書かれた`api`というファイルを作成してください。

### ビルド

```bash
$ make build
```

## インストール

### Windows

生成されるバイナリを配置したいディレクトリに移動させればOKです。
定期的に実行するためタスクスケジューラなどを使ってください。

### Linux

```bash
$ make install
```

## 実行

- 通常実行（グローバルIPアドレスに変更があったらSlackに投げる）

```bash
$ ./IP-notificater
```

タスク例：５分毎に実行


- Cron実行（その日のグローバルIPをSlackに投げる）

```bash
$ ./IP-notificater --cron
```

タスク例：毎日0:00に実行