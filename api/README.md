# Shorten-me-api

## Gerenciador de dependências
Para esse projeto foi utilizado o [Godep](https://github.com/tools/godep) como gerenciador de dependências.

## Variáveis ambientes
Abaixo a lista de variáveis ambientes de configuração.

| Variavel  | Valor Padrão  |
|:-:|:-:|
| PORT | 8081 |
| DB_DRIVER | postgres |
| DB_HOST | localhost |
| DB_NAME | shorten |
| DB_USER | wilson |
| DB_PASSWORD | 1234 |
| DB_SSL_MODE | disable |
| DB_MAX_CONNECTION | 1 |
| DB_LOG_MODE | true |

## Testes
Para todar o test em todos os pacotes use o comando:
```sh
$ go test -v ./...
```

## Build
O build irá gerar os binários nativos para que não seja necessário rodar a aplicação a partir do código fonte, para realizar um novo build vá até a pasta raiz:
```sh
$ go build
```

## Iniciar aplicação
Para iniciar a API apartir do código fonte:
```sh
$ go run server.go
```