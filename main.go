package main

import (
	"log"
	"net/http"
	"os"
	"time"

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

	//Createing custom Server properties to better fine-tune the server details and run away from possible DDOS attacks by setting
	//the timeout and the read and write timeouts
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Serve the server and replace the default http.ListenAndServe() with the custom server
	server.ListenAndServe()

}
