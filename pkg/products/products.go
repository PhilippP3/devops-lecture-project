package products

// Product represents a product in the store
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
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
