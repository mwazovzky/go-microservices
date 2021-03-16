package data

// Product defines the structure for an API product
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

var productRepository = NewRepository()

func GetProducts() []*Product {
	return productRepository.GetProducts()
}

func AddProduct(p Product) {
	productRepository.AddProduct(p)
}

func UpdateProduct(p Product) error {
	return productRepository.UpdateProduct(p)
}

func DeleteProduct(id int) error {
	return productRepository.DeleteProduct(id)
}
