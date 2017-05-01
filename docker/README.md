# Docker Compose
Gerando os containers a partir do `docker-compose.yml`. O arquivo está configurado para gerar 3 container contendo a camada de banco de dados, de servidor API e aplicação web-client todas separadas, pela praticidade de manutenção de escalabilidade.

## Rodando o compose
O comando abaixo já gera as imagens e também os containers. O parametro `-d` ao final é para que o processo não prenda o terminal.
```sh
$ docker-compose up -d
```

Ao final do processo, é necessárior iniciar a API com o seguinte comando:
```sh
$ docker exec shorten-api /bin/./server -d
```

## Testando o client
Para abrir o web-client, basta digitar o seguinte endereço `http://localhost:8080/`.

---

# Dockerfile

## Gerando a imagem
Para gerar a imagem do Docker, vá até a pasta **docker**, abra o terminal e digite:
```sh
$ docker build -t shorten/api .
```

## Criando o container a partir da imagem
Após executar o comando aguarde alguns segundo e verifique se o container foi criado e está executando `docker ps -a`
```sh
$ docker run -d -p 8081:8081 --name shorten-api shorten/api
```

## Executando a API
Após o container ser criado, excute a API com o comando abaixo, em seguida irá mostrar o log de requisições realizadas.
```sh
$ docker exec shorten-api /bin/./server -d
```

---

# Testando com cURL
Para saber mais sobre requisições utilizando a ferramenta cURL, veja este [artigo](http://www.diego-garcia.info/2014/12/13/use-o-curl/) de introdução.

## Resquest POST sem Alias
```sh
$ curl -H "Content-type: application/json" -X POST -d '{"address":"www.google.com.br"}' "http://localhost:8081/api/v1/url/shorten"
```

## Resquest POST com Alias
```sh
$ curl -H "Content-type: application/json" -X POST -d '{"address":"www.google.com.br", "alias": "wilson"}' "http://localhost:8081/api/v1/url/shorten"
```