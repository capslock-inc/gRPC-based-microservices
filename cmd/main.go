package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// handler for root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("visited")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		log.Printf("DATA :: %s", body)
	})

	// handler for page_one
	http.HandleFunc("/page_one", func(w http.ResponseWriter, r *http.Request) {
		log.Print("page_one Visited")
	})

	// listing
	http.ListenAndServe(":9090", nil)
}
