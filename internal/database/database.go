package database

import (
	"database/sql"
	"fmt"
	"log"

	"go-api-starter/internal/config"

	_ "github.com/lib/pq"
)

func Connect(cfg config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar com banco:", err)
	}

	log.Println("Banco conectado com sucesso")

	return db
}
