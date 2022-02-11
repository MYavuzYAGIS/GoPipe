package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})

	http.HandleFunc("/ayla", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello Ayla")
		fmt.Println("Listening the port 9090")
	})
	http.ListenAndServe(":9090", nil)

}
