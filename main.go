package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Data: %s", data)
		fmt.Fprintf(rw, "Hello, %s", data)

	})

	http.ListenAndServe(":9090", nil)

}
