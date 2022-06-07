package app

import (
	"github.com/jmoiron/sqlx"
	"os"
)

//service starting

type ServiceConfig struct {
	address string
	port    string
}

func sanityCheckService() *ServiceConfig {
	var serviceConfig ServiceConfig
	if os.Getenv("SERVER_ADDRESS") == "" {
		serviceConfig.address = "0.0.0.0"
	} else {
		serviceConfig.address = os.Getenv("SERVER_ADDRESS")
	}
	if os.Getenv("SERVER_PORT") == "" {
		serviceConfig.port = "9000"
	} else {
		serviceConfig.port = os.Getenv("SERVER_PORT")
	}
	return &serviceConfig
}

//Database connection starting

type DbConnectionConfig struct {
	DbUser     string
	DbPassword string
	DbAddr     string
	DbPort     string
	DbName     string
}

func SanityCheckDb() *DbConnectionConfig {
	var dbConnectionConfig DbConnectionConfig
	if os.Getenv("DB_USER") == "" {
		dbConnectionConfig.DbUser = "postgres"
	} else {
		dbConnectionConfig.DbUser = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		dbConnectionConfig.DbPassword = "password"
	} else {
		dbConnectionConfig.DbPassword = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("DB_ADDR") == "" {
		dbConnectionConfig.DbAddr = "172.22.2.215"
	} else {
		dbConnectionConfig.DbAddr = os.Getenv("DB_ADDR")
	}
	if os.Getenv("DB_PORT") == "" {
		dbConnectionConfig.DbPort = "5432"
	} else {
		dbConnectionConfig.DbPort = os.Getenv("DB_PORT")
	}
	if os.Getenv("DB_NAME") == "" {
		dbConnectionConfig.DbName = "postgres"
	} else {
		dbConnectionConfig.DbName = os.Getenv("DB_NAME")
	}
	return &dbConnectionConfig
}

func getDbClient() *sqlx.DB {
	dbConnectionConfig := SanityCheckDb()
	dataSourse := "user=" + dbConnectionConfig.DbUser + " password=" + dbConnectionConfig.DbPassword + " host=" + dbConnectionConfig.DbAddr + " port=" + dbConnectionConfig.DbPort + " dbname=" + dbConnectionConfig.DbName + " sslmode=disable"
	pool, err := sqlx.Connect("postgres", dataSourse)
	//pool, err := sqlx.Connect("postgres", "user=postgres password=password host=172.22.2.215 port=5432 dbname=IoT_db_test sslmode=disable")
	if err != nil {
		panic(err)
	}
	return pool
}
