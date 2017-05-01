# Encurtador de URL

Projeto realizado com o intuito validar os conhecimentos na programação e motivação para aprender novas tecnologias.

## Stack Tecnológico
- **PostgreSQL 9.6**
- [**Hashids.org**](http://hashids.org/postgresql/)
- **Golang 1.7**
- **AngularJS v1**
- **Gulp**
- **Docker**

## Executando a aplicação
- [Gerando o deploy da API com Docker](https://github.com/wilsontamarozzi/bemobi-hire-me/tree/master/docker)

## Algoritmo sugerido
O algoritmo escolhido foi o **hashids** que é uma extensão PostgreSQL responsável por gerar um pequeno hash a partir de um inteiro. As duas grandes vantagens de escolher essa API de terceiros é porque ela gerar o hash seguindo uma sequência lógica e não de forma totalmente randômica, e apesar disso ela garante o segredo da sequência a partir de uma palavra segredo que é possível configurar. Ela também consegue se auto administrar em relação ao tamanho do hash gerado, iniciando se com 2 caracteres e se expandindo conforme a necessidade.
O motivo da escolha foi porque apesar desse hash parecer randômico para o usuário final, ele precisa seguir uma sequência lógica para ser gerado, sem isso teríamos os seguintes problemas:

- Caso a aplicação fique responsável por gerar esse hash de forma totalmente randômica, todas vezes que um usuário final solicitar um encurtamento sem passar um alias, a aplicação iria ter que gerar um novo hash randômico que seria necessário fazer um request ao banco de dados para analisar se já existe esse hash salvo, caso já exista, o banco devolveria o retorno para a aplicação e a aplicação teria que refazer todo o processo até encontrar uma que não exista. Pensando que a quantidade de carácter do hash limita a quantidade de combinações, ao chegar próximo ao limite de combinações possíveis, a aplicação passaria muito tempo gerando um novo hash e verificando se já existe no banco. Isso por sua vez iria fazer com que muitas requisições fossem feitas ao banco de dados, pensando que a aplicação irá fornecer serviço para diversos usuários simultâneos isso poderia sobrecarregar a máquina servidor.
- Seguindo a lógica do problema acima, temos outros grande problema de consumo de recursos, onde hoje muitos servidores como Heroku e Amazon possui limites de requests e recursos em seus planos. Como a aplicação teria que gerar um hash novo e verificar se ele já existe a todo momento, o consumo seria elevado e desnecessário, gerando um desperdício de orçamento nos planos.

## Bonus Points
- [x] Crie testcases para todas as funcionalidades criadas
- [x] Crie um endpoint que mostre as dez URL's mais acessadas
- [x] Crie um client para chamar sua API
- [x] Faça um diagrama de sequência da implementação feita nos casos de uso (Dica, use o https://www.websequencediagrams.com/)
- [x] Monte um deploy da sua solução utilizando containers

## Diagramas de Sequência
![](https://github.com/wilsontamarozzi/bemobi-hire-me/blob/master/docs/encurtar%20url.png?raw=true)

![](https://github.com/wilsontamarozzi/bemobi-hire-me/blob/master/docs/recuperar%20url.png?raw=true)
