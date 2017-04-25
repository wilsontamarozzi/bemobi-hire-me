FROM library/postgres
MAINTAINER wilson.tamarozzi@gmail.com

ENV POSTGRES_USER wilson
ENV POSTGRES_PASSWORD 1234
ENV POSTGRES_DB shorten

# cria a pasta onde fica extensão se não houver
RUN /bin/mkdir -p '/usr/lib/postgresql/9.6/lib'
RUN /bin/mkdir -p '/usr/share/postgresql/9.6/extension'

# adicionar os arquivos da extensão
ADD /postgresql-extension_hashids/pg_hashids.so /usr/lib/postgresql/9.6/lib/
ADD /postgresql-extension_hashids/pg_hashids--1.2.sql /usr/share/postgresql/9.6/extension/
ADD /postgresql-extension_hashids/pg_hashids--1.1--1.2.sql /usr/share/postgresql/9.6/extension/
ADD /postgresql-extension_hashids/pg_hashids--1.0--1.1.sql /usr/share/postgresql/9.6/extension/

# adiciona e executa o script do banco
ADD init.sql /docker-entrypoint-initdb.d/
# sobe o programa
ADD /server /home/
# executar o programa
RUN cd /home && ./server &


# RASCUNHO
# atualiza a lista de pacotes
# RUN apt-get update
# instala o pacote dev do postgres
# RUN apt-get install -y postgresql-server-dev-all
# limpa pacotes
# RUN apt-get clean
# sobe a pasta com a extensão hashids do postgres
# ADD /postgresql-extension_hashids.tar.gz /tmp/
# compila extensão hashids do postgres
# RUN cd tmp/postgresql-extension_hashids && make; make install
