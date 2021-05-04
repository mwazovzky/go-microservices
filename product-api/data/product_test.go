package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "Alex",
		Price: 1.00,
		SKU:   "aaa-bbb-ccc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
