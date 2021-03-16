package data

import (
	"fmt"
	"time"
)

// ErrProductNotFound is an error raised when a product can not be found in the repository
var ErrProductNotFound = fmt.Errorf("Product not found")

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAT:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fdj347",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAT:   time.Now().UTC().String(),
	},
}

type Repository struct {
	index []*Product
}

// NewValidation creates a new Validation type
func NewRepository() *Repository {
	return &Repository{productList}
}

// GetProducts returns all products from the repository
func (pr *Repository) GetProducts() []*Product {
	return pr.index
}

// FindProduct returns a single product which matches the id from the repository.
// If a product is not found this function returns a ProductNotFound error
func (pr *Repository) FindProduct(id int) (*Product, error) {
	i := pr.findIndex(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

// AddProduct adds a new product to the repository
func (pr *Repository) AddProduct(p Product) {
	maxID := pr.index[len(pr.index)-1].ID
	p.ID = maxID + 1
	pr.index = append(pr.index, &p)
}

// UpdateProduct replaces a product in the repository with the given item.
// If a product with the given id does not exist in the repository
// this function returns a ProductNotFound error
func (pr *Repository) UpdateProduct(p Product) error {
	i := pr.findIndex(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}

	pr.index[i] = &p

	return nil
}

// DeleteProduct deletes a product from the repository
func (pr *Repository) DeleteProduct(id int) error {
	i := pr.findIndex(id)
	if i == -1 {
		return ErrProductNotFound
	}

	pr.index = append(pr.index[:i], pr.index[i+1:]...)

	return nil
}

// findIndex finds the index of a product in the repository
// returns -1 when no product can be found
func (pr *Repository) findIndex(id int) int {
	for i, p := range pr.index {
		if p.ID == id {
			return i
		}
	}

	return -1
}
