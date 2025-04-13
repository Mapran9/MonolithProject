package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() error {
	var err error
	dsn := "mono_user:mono1234@tcp(db:3306)/mono_db"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database ping error:", err)
	}

	fmt.Println("Database connection established")
	return nil
}

func RunMigrations(db *sql.DB) error {
	migrationsPath := "db/migrations/"
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(migrationsPath, file.Name())

		sqlFile, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %v", filePath, err)
		}

		_, err = db.Exec(string(sqlFile))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %v", filePath, err)
		}

		fmt.Printf("Executed migration: %s\n", filePath)
	}

	return nil
}
