-- cria usuário para o app
--CREATE USER wilson WITH PASSWORD '1234';
-- cria o banco de dados
--CREATE DATABASE shorten;
-- da todos previlegios para o usuário no banco
--GRANT ALL PRIVILEGES ON DATABASE shorten TO wilson;
-- loga no novo banco
\c shorten;
-- adiciona as extensões ao banco
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_hashids";
-- cria o banco de dados
CREATE TABLE "urls" ("uuid" uuid UNIQUE DEFAULT uuid_generate_v4(),"serial" serial UNIQUE,"address" varchar(500) NOT NULL,"alias" varchar(30) UNIQUE,"view" integer , PRIMARY KEY ("uuid","serial"));
-- cria função que analisa e gera o alias
CREATE OR REPLACE FUNCTION urls_pre_insert() RETURNS TRIGGER AS $$
BEGIN
	IF (NEW.alias IS NULL OR NEW.alias = "") THEN
    	NEW.alias := hash_encode(NEW.serial, "secret_salt", 1);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- criar o gatilho que chama a função antes do insert
CREATE TRIGGER urls_pre_insert BEFORE INSERT ON urls FOR EACH ROW EXECUTE PROCEDURE urls_pre_insert();