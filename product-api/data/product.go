package data

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`

	// Date product was created at
	CreatedAt string `json:"-"`

	// Date product was updated at
	UpdatedAT string `json:"-"`
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
