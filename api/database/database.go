package database

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wilsontamarozzi/bemobi-hire-me/api/models"
	"log"
	"os"
	"sync"
	"fmt"
)

const (
	ENV_DB_DRIVER         = "DB_DRIVER"
	ENV_DB_HOST           = "DB_HOST"
	ENV_DB_NAME           = "DB_NAME"
	ENV_DB_USER           = "DB_USER"
	ENV_DB_PASSWORD       = "DB_PASSWORD"
	ENV_DB_SSL_MODE       = "DB_SSL_MODE"
	ENV_DB_MAX_CONNECTION = "DB_MAX_CONNECTION"
	ENV_DB_LOG_MODE       = "DB_LOG_MODE"
)

var (
	DB_DRIVER         string = "postgres"
	DB_HOST           string = "localhost"
	DB_NAME           string = "shorten"
	DB_USER           string = "wilson"
	DB_PASSWORD       string = "1234"
	DB_SSL_MODE       string = "disable" // disable | require
	DB_MAX_CONNECTION int    = 1
	DB_LOG_MODE       bool   = true
)

var once sync.Once
var DBSession *gorm.DB

func init() {
	getEnvDatabaseConfig()
	GetInstance()
	//RebuildDataBase()
	//TruncateTablesDB()
}

func getEnvDatabaseConfig() {
	dbDriver := os.Getenv(ENV_DB_DRIVER)
	dbHost := os.Getenv(ENV_DB_HOST)
	dbName := os.Getenv(ENV_DB_NAME)
	dbUser := os.Getenv(ENV_DB_USER)
	dbPassword := os.Getenv(ENV_DB_PASSWORD)
	dbSslMode := os.Getenv(ENV_DB_SSL_MODE)
	dbMaxConnection := os.Getenv(ENV_DB_MAX_CONNECTION)
	dbLogMode := os.Getenv(ENV_DB_LOG_MODE)

	maxConnection, err1 := govalidator.ToInt(dbMaxConnection)
	logMode, err2 := govalidator.ToBoolean(dbLogMode)

	if len(dbDriver) > 0 {
		DB_DRIVER = dbDriver
	}
	if len(dbHost) > 0 {
		DB_HOST = dbHost
	}
	if len(dbName) > 0 {
		DB_NAME = dbName
	}
	if len(dbUser) > 0 {
		DB_USER = dbUser
	}
	if len(dbPassword) > 0 {
		DB_PASSWORD = dbPassword
	}
	if len(dbSslMode) > 0 {
		DB_SSL_MODE = dbSslMode
	}
	if err1 == nil {
		DB_MAX_CONNECTION = int(maxConnection)
	}
	if err2 == nil {
		DB_LOG_MODE = logMode
	}
}

func GetInstance() *gorm.DB {
	once.Do(func() {
		DBSession = buildConnection()
	})

	return DBSession
}

func buildConnection() *gorm.DB {
	strConnection := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_SSL_MODE, DB_PASSWORD)
	db, err := gorm.Open(DB_DRIVER, strConnection)
	if err != nil {
		log.Print("Não foi possivel conectar ao banco de dados")
		panic(err)
	}

	log.Print("Conexão com o banco realizada com sucesso")
	//Ativa log de todas as saidas da conexão (SQL)
	db.LogMode(DB_LOG_MODE)
	//Seta o maximo de conexões
	db.DB().SetMaxIdleConns(DB_MAX_CONNECTION)
	db.DB().SetMaxOpenConns(DB_MAX_CONNECTION)

	return db
}

func RebuildDataBase() {
	CreateExtensionsPostgreSQL()
	DropTriggersIfExists()
	DropFunctionsIfExists()
	DropTablesIfExists()
	AutoMigrate()
	CreateFunctionsDB()
	CreateTriggersDB()
}

func DropTablesIfExists() {
	GetInstance().Exec("DROP TABLE IF EXISTS urls CASCADE;")
}

func TruncateTablesDB() {
	GetInstance().Exec("TRUNCATE TABLE urls;")
}

func AutoMigrate() {
	GetInstance().AutoMigrate(&models.URL{})
}

func CreateFunctionsDB() {
	GetInstance().Exec(`
		CREATE OR REPLACE FUNCTION urls_pre_insert() RETURNS TRIGGER AS $$
		BEGIN
			IF (NEW.alias IS NULL OR NEW.alias = '') THEN
		    	NEW.alias := hash_encode(NEW.serial, 'secret_salt', 1);
		    END IF;
		    RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;`)
}

func DropFunctionsIfExists() {
	GetInstance().Exec(`DROP FUNCTION IF EXISTS urls_pre_insert();`)
}

func CreateTriggersDB() {
	GetInstance().Exec(`CREATE TRIGGER urls_pre_insert BEFORE INSERT ON urls FOR EACH ROW EXECUTE PROCEDURE urls_pre_insert();`)
}

func DropTriggersIfExists() {
	GetInstance().Exec(`DROP TRIGGER IF EXISTS urls_pre_insert ON urls;`)
}

func CreateExtensionsPostgreSQL() {
	GetInstance().Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	GetInstance().Exec(`CREATE EXTENSION IF NOT EXISTS "pg_hashids";`)
}
