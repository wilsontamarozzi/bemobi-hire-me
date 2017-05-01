FROM library/postgres
MAINTAINER Wilson Tamarozzi <wilson.tamarozzi@gmail.com>

# 'Criando as pastas onde irá ficar a extensão pg_hashids'
RUN mkdir -p '/usr/lib/postgresql/9.6/lib' \
mkdir -p '/usr/share/postgresql/9.6/extension'

# 'Adicionando os arquivos da extensão pg_hashids'
ADD /pg_hashids/pg_hashids.so /usr/lib/postgresql/9.6/lib/
ADD ["/pg_hashids/pg_hashids.control", "/pg_hashids/pg_hashids--1.2.sql", "/pg_hashids/pg_hashids--1.1--1.2.sql", "/pg_hashids/pg_hashids--1.0--1.1.sql", "/usr/share/postgresql/9.6/extension/"]

# 'Adicionando o script do banco de dados'
ADD init.sql /docker-entrypoint-initdb.d/