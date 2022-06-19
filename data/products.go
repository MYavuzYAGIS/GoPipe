package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"` // internal product identifier
	CreatedOn   string  `json:"-"`   // will not be encoded into JSON
	UpdatedOn   string  `json:"-"`   // will not be encoded into JSON
	DeletedOn   string  `json:"-"`   // will not be encoded into JSON
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil

}

var ErrorProductNotFound = fmt.Errorf("Prod not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "lorem",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Ayla",
		Description: "Frothy azbla coffee",
		Price:       2.45,
		SKU:         "ipsum",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          3,
		Name:        "Espresso",
		Description: "Frothy ayla coffee",
		Price:       2.45,
		SKU:         "ipsm",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
