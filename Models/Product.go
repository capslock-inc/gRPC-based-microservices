package models

import (
	"encoding/json"
	"io"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Allproducts []*Product

func (p *Allproducts) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Allproducts {
	return Productlist
}

var Productlist = []*Product{
	&Product{
		Id:          1001,
		Name:        "Capichino",
		Description: "nothing special",
		Price:       2.45,
		SKU:         "Abc",
		CreatedOn:   "today",
		UpdatedOn:   "today",
		DeletedOn:   "none",
	},
	&Product{
		Id:          1002,
		Name:        "latte",
		Description: "nothing special",
		Price:       2.22,
		SKU:         "Abc",
		CreatedOn:   "today",
		UpdatedOn:   "today",
		DeletedOn:   "none",
	},
}
