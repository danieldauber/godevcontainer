package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello WOrlds")
	})

	fmt.Println(http.ListenAndServe(":8083", nil))
}
