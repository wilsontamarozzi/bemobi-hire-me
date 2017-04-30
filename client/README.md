### NPM
Baixando as bibliotecas utilziadas.
```
$ npm install
```

### Build
Gerando a pasta de distribuição pronta para ser hospedada.

```
$ gulp build
```

### Gulp
Toda a estrutura de configuração das tarefas do Gulp estão no arquivo ```config.js```.

```
Caminho: gulp/config.js
```

Lista de todas as tarefas prontas no Gulp.

- ```Clean```: Limpa a pasta de destribuição e excluir o arquivo zip da extensão.
- ```Compress-css```: Unifica e comprime todos os arquivos css para ```styles.min.css```.
- ```Compress-js```: Unifica e comprime todos os arquivos js para ```all.min.js```.
- ```Concat-lib```: Unifica todos os arquivos js das bibliotecas já minificadas para ```lib.min.js```.
- ```Copy-other-files```: Copia todos os arquivos fonts que não precisam passar por tratamente para a pasta de distribuição (imagens, html, fonts boostrap e o index de produção).
- ```Jshint```: Escaneia os arquivos js a procura de erros de syntax.
- ```Pack```: Realiza a compreensão da pasta de distribuição para ```zipped_dist.zip```.
- ```Build```: Função responsável por compilar todas funções de uma vez e gerar a pasta de distribuição.