package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func gormConnect() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	// .envを取得、代入
	DBMS := os.Getenv("DBMS_NAME")
	USER := os.Getenv("USER_NAME")
	PASS := os.Getenv("PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := gormConnect()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// // Read
	var product Product
	first_row := db.First(&product, 1) // find product with integer primary key
	fmt.Println(first_row)
	first_row = db.First(&product, "code = ?", "D42") // find product with code D42
	fmt.Println(first_row)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&product, 1)
}
