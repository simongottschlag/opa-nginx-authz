package main

import (
	"fmt"
	"net/http"
)

func echo(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%#v\n", req)
	fmt.Fprintf(w, "%#v", req)
}

func main() {
	http.HandleFunc("/", echo)
	http.ListenAndServe(":8081", nil)
}
