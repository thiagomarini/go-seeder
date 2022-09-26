// pkg/seeds/seeds.go
package seeds

import (
	"fmt"
	"gorm.io/gorm"
	"seeder/seed"
)

func All() []seed.Seed {
	return []seed.Seed{
		seed.Seed{
			Name: "CreateHeadphone",
			Run: func(db *gorm.DB) error {
				fmt.Println("Creating a headphone...")
				return CreateProduct(db, "HF", 30)
			},
		},
	}
}
