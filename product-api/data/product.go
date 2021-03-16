package data

import (
	"fmt"
	"time"
)

// https://github.com/go-playground/validator
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedAt   string  `json:"-"`
	UpdatedAT   string  `json:"-"`
	DeletedAt   string  `json:"-"`
}

func (p *Product) Validate() error {
	validator := NewValidator()
	return validator.validate.Struct(p)
}

type Products []*Product

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAT:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fdj347",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAT:   time.Now().UTC().String(),
	},
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, index, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[index] = p
	return nil
}

func DeleteProduct(id int) error {
	_, index, err := findProduct(id)
	if err != nil {
		return err
	}

	productList = append(productList[:index], productList[index+1])

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for index, p := range productList {
		if p.ID == id {
			return p, index, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	last := productList[len(productList)-1]
	return last.ID + 1
}
