# go-chat
## Description
goによるチャットアプリ用のapi server

## 環境構築
### ローカル環境の構築
1. Dockerとdocker-composeのインストール
- [Docker のインストール(公式ドキュメント)](https://docs.docker.jp/engine/installation/index.html)
- [Docker Compose のインストール(公式ドキュメント)](https://docs.docker.jp/compose/install.html)
2. イメージビルド
```
docker-compose build
```
3. パッケージインストール
```
docker-compose run --rm app make install
docker-compose run --rm node yarn install
```
4. migration
```
docker-compose run --rm app sql-migrate up
```
4. コンテナ作成
```
docker-compose up -d
```
5. コンテナの起動確認
```
docker-compose ps
```
### 本番環境用のDocker imageのビルド
```
docker build -t go-chat .
docker run --name go-chat -d -it -p 8080:8080  go-chat
```

