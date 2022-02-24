package data

import "time"

type Product struct {
	ID int 
	Name string
	Description string
	Price float32
	SKU string  // internal product identifier
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

var productList =[]*Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "lorem",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Ayla",
		Description: "Frothy azbla coffee",
		Price: 2.45,
		SKU: "ipsum",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},	
	&Product{
		ID: 3,
		Name: "Espresso",
		Description: "Frothy ayla coffee",
		Price: 2.45,
		SKU: "ipsm",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}

