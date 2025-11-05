package database

import (
	"backend-inventory-go/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	connectionString := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s;encrypt=true;trustServerCertificate=true",
		cfg.Server, cfg.Port, cfg.Database, cfg.User, cfg.Password)

	DB, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		return fmt.Errorf("error al abrir la conexión: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error al conectar con la base de datos: %v", err)
	}

	log.Println("✅ Conexión a SQL Server establecida correctamente")
	return nil
}

func GetDB() *sql.DB {
	return DB
}
