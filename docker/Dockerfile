FROM postgres
MAINTAINER wilson.tamarozzi@gmail.com
# atualiza a lista de pacotes
RUN apt-get update
# instala o pacote dev do postgres
RUN apt-get install -y postgresql-server-dev-all
# limpa pacotes
RUN apt-get clean
# sobe a pasta com a extensão hashids do postgres
ADD /postgresql-extension_hashids.tar.gz /tmp/
# compila extensão hashids do postgres
RUN cd tmp/postgresql-extension_hashids && make; make install
# cria o usuário e banco de dados
ADD init.sql /docker-entrypoint-initdb.d/
# sobe o programa
ADD /server /home/
# executar o programa
RUN cd /home && ./server &