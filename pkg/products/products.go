package models

// Product represents a product in the store
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Static data for three products
var products = []Product{
	{ID: 1, Name: "Samsung Galaxy A 35", Price: 309.99},
	{ID: 2, Name: "Xiaomi Redmi Note 14", Price: 231.21},
	{ID: 3, Name: "Honor Magic7 Pro", Price: 1379.99},
}

// FindProductByID searches for a product by its ID
func FindProductByID(products []Product, id int) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
