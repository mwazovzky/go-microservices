package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello World!")
}

func main() {
	// curl -v localhost:9090 -d "mike"
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Something went wrong"))
			http.Error(rw, "Something went wrong", http.StatusBadRequest)
			log.Printf("Bad request: %s", err)
			return
		}

		log.Printf("Data: %s", data)
		fmt.Fprintf(rw, "Hello %s!\n", data)
	})

	http.HandleFunc("/goodby", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodby World!")
	})

	log.Println("Starting http server at :9090")
	http.ListenAndServe(":9090", nil)
}
