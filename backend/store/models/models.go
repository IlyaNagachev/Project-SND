package models

import "github.com/scylladb/gocqlx/v3/table"

const ProductTable = "products"

type Tables struct {
	ProductTable *table.Table
}

type Products []*Product

type Product struct{
	Uuid string `gorm:"uuid" json:"uuid"`
	Class string `db:"class" json:"class"`
	Name string `db:"name" json:"name"`
	Price float32 `db:"price" json:"price"`
	Img string `db:"img" json:"img"`
	Description string `db:"description" json:"description"`
	Specs map[string]string `db:"specs" json:"specs"`
}
