package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/", Hello)
	fmt.Println(http.ListenAndServe(port, nil))

}

func Hello(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(rw, "Hello")
}
