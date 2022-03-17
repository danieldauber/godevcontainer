package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", Hello)
	fmt.Println(http.ListenAndServe(":80", nil))

}

func Hello(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(rw, "Hello")
}
