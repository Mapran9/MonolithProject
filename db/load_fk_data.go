package db

import (
	"database/sql"
	"fmt"
	"os"
)

func UpdateForeignKeys(db *sql.DB) error {
	sqlFiles := []string{
		"db/relations/001_update_cart.sql",
		"db/relations/002_update_order.sql",
		"db/relations/003_update_payment.sql",
	}

	for _, file := range sqlFiles {
		// read SQL
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("error reading SQL file %s: %v", file, err)
		}
		sqlCommands := string(sqlBytes)

		// exec SQL
		_, err = db.Exec(sqlCommands)
		if err != nil {
			return fmt.Errorf("error executing SQL from %s: %v", file, err)
		}

		fmt.Printf("Executed %s successfully!\n", file)
	}

	return nil
}
