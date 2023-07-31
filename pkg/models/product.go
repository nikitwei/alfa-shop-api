package models

import (
	"github.com/nikitwei/alfa-shop-api/pkg/config"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type MsProduct struct {
	ID          int64     `gorm:"primaryKey;column:id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	PriceUnit   string    `gorm:"column:priceUnit"`
	CreatedAt   time.Time `gorm:"column:createdAt"`
	CreatedBy   int64     `gorm:"column:createdBy"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
	UpdatedBy   int64     `gorm:"column:updatedBy"`
	gorm.Model
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&MsProduct{})
}

func AddProduct(p *MsProduct) *MsProduct {
	p.CreatedBy = 1
	p.UpdatedBy = 1
	db.Create(&p)
	return p
}

func GetProducts() []MsProduct {
	var products []MsProduct
	db.Find(&products)
	return products
}

func GetProductID(Id int64) (*MsProduct, *gorm.DB) {
	var product MsProduct
	db := db.Where("ID=?", Id).Find(&product)
	return &product, db
}

func DeleteProduct(Id int64) MsProduct {
	var product MsProduct
	db.Where("ID=?", Id).Delete(&product)
	return product
}
