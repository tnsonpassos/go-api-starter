package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"

	"go-api-starter/internal/config"
	"go-api-starter/internal/database"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Use: go run cmd/migrate/main.go up")
	}

	direction := os.Args[1]

	cfg := config.Load()

	db := database.Connect(cfg)
	defer db.Close()

	ensureSchemaMigrationsTable(db)

	switch direction {
	case "up":
		runUpMigrations(db)
	case "down":
		runDownMigration(db)
	default:
		log.Fatalf("Comando inválido: %s. Use: up ou down", direction)
	}
}

func ensureSchemaMigrationsTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		log.Fatal("Erro ao criar tabela schema_migrations:", err)
	}
}

func runUpMigrations(db *sql.DB) {
	files, err := filepath.Glob("migrations/*.up.sql")
	if err != nil {
		log.Fatal("Erro ao buscar migrations:", err)
	}

	if len(files) == 0 {
		log.Println("Nenhuma migration encontrada")
		return
	}

	for _, file := range files {
		version := getMigrationVersion(file)

		applied, err := isMigrationApplied(db, version)
		if err != nil {
			log.Fatal("Erro ao verificar migration:", err)
		}

		if applied {
			log.Println("Migration já aplicada:", version)
			continue
		}

		sqlContent, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("Erro ao ler migration:", err)
		}

		_, err = db.Exec(string(sqlContent))
		if err != nil {
			log.Fatalf("Erro ao executar migration %s: %v", version, err)
		}

		err = registerMigration(db, version)
		if err != nil {
			log.Fatalf("Erro ao registrar migration %s: %v", version, err)
		}

		log.Println("Migration aplicada:", version)
	}
}

func isMigrationApplied(db *sql.DB, version string) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM schema_migrations
			WHERE version = $1
		)
	`

	err := db.QueryRow(query, version).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func registerMigration(db *sql.DB, version string) error {
	_, err := db.Exec(`
		INSERT INTO schema_migrations (version)
		VALUES ($1)
	`, version)

	return err
}

func getMigrationVersion(path string) string {
	filename := filepath.Base(path)

	version := strings.TrimSuffix(filename, ".up.sql")

	return version
}

func runDownMigration(db *sql.DB) {
	version, err := getLastAppliedMigration(db)
	if err != nil {
		log.Fatal("Erro ao buscar última migration:", err)
	}

	if version == "" {
		log.Println("Nenhuma migration aplicada para reverter")
		return
	}

	file := filepath.Join("migrations", version+".down.sql")

	sqlContent, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Erro ao ler migration down %s: %v", version, err)
	}

	_, err = db.Exec(string(sqlContent))
	if err != nil {
		log.Fatalf("Erro ao executar migration down %s: %v", version, err)
	}

	err = unregisterMigration(db, version)
	if err != nil {
		log.Fatalf("Erro ao remover registro da migration %s: %v", version, err)
	}

	log.Println("Migration revertida:", version)
}

func getLastAppliedMigration(db *sql.DB) (string, error) {
	var version string

	err := db.QueryRow(`
		SELECT version
		FROM schema_migrations
		ORDER BY applied_at DESC
		LIMIT 1
	`).Scan(&version)

	if err == sql.ErrNoRows {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return version, nil
}

func unregisterMigration(db *sql.DB, version string) error {
	_, err := db.Exec(`
		DELETE FROM schema_migrations
		WHERE version = $1
	`, version)

	return err
}
