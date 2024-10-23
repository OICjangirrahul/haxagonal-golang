# ユーザーポータル API

このプロジェクトは、ユーザーアカウントの管理を行うための API です。ユーザーの登録、更新、取得が可能です。

## セットアップ

このプロジェクトをローカル環境でセットアップする手順は以下の通りです。

### 前提条件

- [Docker](https://www.docker.com/get-started) がインストールされていること
- [Docker Compose](https://docs.docker.com/compose/install/) がインストールされていること

### セットアップ対応

```
1. step
cp .env.sample .env
必要な設定url, portなど書き換え

2. step
docker-compose build

```

### swagger の作成手順

```
docker exec -it portal-api swag init -g cmd/main.go
docker exec -it portal-api golangci-lint run
```
