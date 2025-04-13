package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Load CSV Data to Table
func LoadCSVToDB(tableName, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) < 2 {
		return fmt.Errorf("no data found in %s", filePath)
	}

	columns := strings.Join(records[0], ",")
	placeholders := strings.Repeat("?,", len(records[0])-1) + "?"

	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, columns, placeholders))
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, row := range records[1:] {
		values := make([]interface{}, len(row))
		for i, v := range row {
			values[i] = v
		}
		_, err = stmt.Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("Data loaded into %s successfully!\n", tableName)
	return nil
}
