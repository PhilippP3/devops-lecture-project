package models

// Product represents a product in the store
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Static data for three products
var products = []Product{
	{ID: 1, Name: "Product 1", Price: 19.99},
	{ID: 2, Name: "Product 2", Price: 29.99},
	{ID: 3, Name: "Product 3", Price: 39.99},
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
