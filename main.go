package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// return 404 on / path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		file, err := ioutil.ReadFile("./assets/404.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatal("Can't find error html page")
		}
		w.Write(file)
	})

	// return 200 on healthz path
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthy!"))
	})

	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}
