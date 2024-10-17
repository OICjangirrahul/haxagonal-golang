# ユーザーポータル API

このプロジェクトは、ユーザーアカウントの管理を行うための API です。ユーザーの登録、更新、取得が可能です。

## セットアップ

このプロジェクトをローカル環境でセットアップする手順は以下の通りです。

### 前提条件

- [Go](https://golang.org/dl/) がインストールされていること
- [Docker](https://www.docker.com/get-started) がインストールされていること
- [Docker Compose](https://docs.docker.com/compose/install/) がインストールされていること

### インストール手順

1. step to run project

```
cp config/configEx.yml config/config.yml
docker-compose up
docker exec -it portal-api go mod tidy
```
