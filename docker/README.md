# Dockerfile

## Gerando a imagem
Para gerar a imagem do Docker, vá até a pasta **docker**, abra o terminal e digite:
```sh
$ docker build -t bemobi/hire-me .
```

## Criando o container a partir da imagem
Após executar o comando aguarde alguns segundo e verifique se o container foi criado e está executando `docker ps -a`
```sh
$ docker run -d -p 8080:8080 --name hire-me-server bemobi/hire-me
```

## Executando a API
Após o container ser criado, excute a API com o comando abaixo, em seguida irá mostrar o log de requisições realizadas.
```sh
$ docker exec -it hire-me-server /home/./server
```

---

# Testando com cURL
Para saber mais sobre requisições utilizando a ferramenta cURL, veja este [artigo](http://www.diego-garcia.info/2014/12/13/use-o-curl/) de introdução.

## Resquest POST sem Alias
```sh
$ curl -H "Content-type: application/json" -X POST -d '{"address":"www.google.com.br"}' "http://localhost:8080/api/v1/url/shorten"
```

## Resquest POST com Alias
```sh
$ curl -H "Content-type: application/json" -X POST -d '{"address":"www.google.com.br", "alias": "wilson"}' "http://localhost:8080/api/v1/url/shorten"
```