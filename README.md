# Conflr

Confluence APIを叩いて取得したSolrを

```sh

.
|-- README.md
|-- app                   # インデクシングCLIのソース
|-- docker-compose.yaml
|-- front                 # フロントエンドのソース
`-- solr
    |-- data              # Solrデータの永続化
    |-- managed-schema
    `-- web.xml

11 directories, 11 files
```

## 使い方

### インデクシング

```sh
$ docker-compose up -d
$ docker-compose exec solr bin/solr create_core -c confl
$ cp ./solr/managed-schema ./solr/data/confl/conf/managed-schema
# localhost:8983 のAdmin画面でコアをリロード
$ cd ./app
$ direnv allow
$ go run *.go import -u <コンフルのユーザ名> -p <コンフルのパスワード>
```

### フロント画面

```sh
$ ./front
$ yarn start
# localhost:3000 に検索画面がある
```
