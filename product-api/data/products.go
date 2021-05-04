package data

import (
	"context"
	"fmt"

	protos "github.com/mwazovzky/microservices-introduction/currency/protos/currency"
)

type Products struct {
	pr *Repository
	cc protos.CurrencyClient
}

func NewProducts(pr *Repository, cc protos.CurrencyClient) *Products {
	return &Products{pr, cc}
}

func (p *Products) GetProducts(currency string) (ProductList, error) {
	productList := p.pr.GetProducts()

	if currency == "" {
		return productList, nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		fmt.Println("[ERROR] fail to get exchange rate", err)
		return nil, err
	}

	list := ProductList{}
	for _, p := range productList {
		np := *p
		np.Price = np.Price * rate
		list = append(list, &np)
	}

	return list, nil
}

func (p *Products) FindProduct(id int, currency string) (*Product, error) {
	product, err := p.pr.FindProduct(id)

	if currency == "" {
		return product, nil
	}

	rate, err := p.getRate(currency)
	if err != nil {
		fmt.Println("[ERROR] fail to get exchange rate", err)
		return nil, err
	}

	// convert product price to specified currency
	np := *product
	np.Price = product.Price * rate

	return &np, err
}

func (p *Products) AddProduct(product Product) {
	p.pr.AddProduct(product)
}

func (p *Products) UpdateProduct(prodyct Product) error {
	return p.pr.UpdateProduct(prodyct)
}

func (p *Products) DeleteProduct(id int) error {
	return p.pr.DeleteProduct(id)
}

func (p *Products) getRate(destination string) (float64, error) {
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["EUR"]),
		Destination: protos.Currencies(protos.Currencies_value[destination]),
	}

	res, err := p.cc.GetRate(context.Background(), rr)

	return res.Rate, err
}
