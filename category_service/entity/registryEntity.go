package entity

import (
	"time"
)

type RegistryEntity struct {
	ID             int64     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Date           time.Time `gorm:"column:date"`
	Category       string    `gorm:"type:varchar(255);column:category"`
	Title          string    `gorm:"type:varchar(255);not null;column:title"`
	Amount         float64   `gorm:"type:varchar(255);not null;column:amount"`
	MyCategory     string    `gorm:"type:varchar(255);column:my_category"`
	FamilyCategory string    `gorm:"type:varchar(255);column:family_category"`
}

func (RegistryEntity) TableName() string {
	return "registry"
}
