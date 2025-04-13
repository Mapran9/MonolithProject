package main

import (
	"github.com/Mapran9/MonolithProject/db"
	"log"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.DB.Close()

	// Step 1: รัน migrations (สร้างตาราง)
	err = db.RunMigrations(db.DB)
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	log.Println("Migrations completed successfully!")

	// Step 2: โหลดข้อมูลจาก CSV
	LoadAndProcessData()

	log.Println("Setup complete!")
}

func LoadAndProcessData() {
	files := map[string]string{
		"products":  "docker/init/product.csv",
		"customers": "docker/init/customer.csv",
		"carts":     "docker/init/cart.csv",
		"orders":    "docker/init/order.csv",
		"payments":  "docker/init/payment.csv",
	}

	for table, filePath := range files {
		err := db.LoadCSVToDB(table, filePath)
		if err != nil {
			log.Fatalf("Error loading %s: %v", table, err)
		}
	}

	log.Println("All CSV data loaded successfully!")

	// อัปเดต Foreign Keys
	err := db.UpdateForeignKeys(db.DB)
	if err != nil {
		log.Fatalf("Error updating foreign keys: %v", err)
	}

	log.Println("All foreign keys updated successfully!")
}
