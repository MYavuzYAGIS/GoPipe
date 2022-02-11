package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
		Addr:         ":9090", // binding all connections to port 9090
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Serve the server and replace the default http.ListenAndServe() with the custom server

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)
	sig := <-sigChannel
	l.Println("Received terminate, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext) // for graceful shutdown

}
