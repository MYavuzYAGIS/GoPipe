package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/MYavuzYAGIS/GoPipe/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		//expect the id in the uri
		path := r.URL.Path
		reg := regexp.MustCompile(`/(\d+)`)
		group := reg.FindAllStringSubmatch(path, -1)
		if len(group) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(group[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.l.Println("Got Id", id)
		p.updateProducts(id, rw, r)
		return

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Oops! Something went wrong", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Products")

	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}
	p.l.Printf("Product: %#v", product)
	data.AddProduct(product)
}

func (p Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle  Complete update of the  product")
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Prod not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Prod not found", http.StatusInternalServerError)
		return
	}

}
