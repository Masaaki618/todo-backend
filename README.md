# todo-backend

## Docker での開発

ホットリロード付きの API と PostgreSQL を起動する。

```bash
docker compose up --build
```

起動後、API は `http://localhost:8080` で待ち受ける。
PostgreSQL は `localhost:5433` で待ち受ける。

このディレクトリ配下の Go ファイルを編集すると、コンテナ内の `air` が変更を検知して自動で再ビルド・再起動する。

アプリコンテナには次の接続文字列が渡される。

```bash
postgres://todo:todo@db:5432/todo?sslmode=disable
```

停止するときは次を実行する。

```bash
docker compose down
```

DB データも削除して初期化したい場合は次を実行する。

```bash
docker compose down -v
```

## OpenAPI

API 仕様は [openapi.yaml](/Users/poi/work/todo/todo-backend/openapi.yaml) に定義している。
