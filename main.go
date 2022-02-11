package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MYavuzYAGIS/GoPipe/handlers"
)

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// })
	l := log.New(os.Stdout, "product-API", log.LstdFlags)

	//Handler defs
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)
	// Mux defs
	serveMux := http.NewServeMux()
	// Mux implementation
	serveMux.Handle("/hello", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)
	// Serve
	http.ListenAndServe(":9090", serveMux)

}
