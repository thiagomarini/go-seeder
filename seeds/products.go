package seeds

import (
	"fmt"
	"gorm.io/gorm"
	"seeder/models"
)

func CreateProduct(db *gorm.DB, code string, price int) error {
	fmt.Println("Creating a product from seeder")
	return db.Create(&models.Product{Code: "D42", Price: 100}).Error
}
