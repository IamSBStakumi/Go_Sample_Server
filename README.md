# Go_Sample_Server

![Go](https://img.shields.io/badge/-Go-F2C63C.svg?logo=go&style=for-the-badge)
![OpenAPI](https://img.shields.io/badge/-OpenAPI-22C535.svg?logo=OpenAPIInitiative&style=for-the-badge)
![YAML](https://img.shields.io/badge/-yaml-337BDD.svg?logo=yaml&style=for-the-badge)

## Install oapi-codegen

[oapi-codegen GitHub](https://github.com/deepmap/oapi-codegen/tree/master)

```bash
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
```

```bash
oapi-codegen -version
```

### Install Golang

- [Windows](./Install_Go/Windows.md)
- [Mac](./Install_Go/Mac.md)
- [Ubuntu](./Install_Go/Ubuntu.md)

## Generate Code

作成されていなければ、config.yaml を作成

記法例

```config.yaml
# oapi-codegenコマンドを使用するための設定ファイル
package: api
output: ./generated/openapi.gen.go # generatedディレクトリが存在しないと失敗する
generate:
  echo-server: true
  models: true
output-options:
  skip-prune: true
```

components ディレクトリが存在するなら、merge OpenAPI yaml の項を参照

oapi-codegen を実行する

```bash
oapi-codegen -config config.yaml openapi.yaml
```

## merge OpenAPI yaml

yaml ファイルを components フォルダに分割していると、`oapi-codegen`実行時にエラーが生じる

そのため、自動で openapi.yaml を統合してくれる`@redocly/cli`を使用する

```bash
npx @redocly/cli bundle {メインのyamlファイル} -o ./generated/allinone.gen.yaml
```

結合後、oapi-codegen を実行する際に結合したファイルを使用する

```bash
oapi-codegen -config config.yaml generated/allinone.gen.yaml
```

# init go server

```bash
go mod init
```

or

```bash
go mod init example.com/m/v2
go mod tidy
```
