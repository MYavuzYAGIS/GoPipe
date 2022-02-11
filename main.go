package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// })
	l := log.New(os.Stdout, "product-API", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	serveMux := http.NewServeMux()
	serveMux.Handle("/hello", helloHandler)
	http.ListenAndServe(":9090", serveMux)

}
