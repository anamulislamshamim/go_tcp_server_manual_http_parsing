package models

// Product represents a simple ecommerce product
type Product struct {
	ID    int    `json:"id"`   // unique indentifier
	Name  string `json:"name"` // Product name
	Price int    `json:"price"`
}

// Inmemory Store acts as our "database"
var Products = []Product{}

// NextId keeps track of the next available ID
var NextId = 1
