package main

import (
	"fmt"
	"net/http"
)

func debug(w http.ResponseWriter, req *http.Request) {

	//fmt.Println("server: req", req, "\nheader", req.Header, "\nbody", req.Body)
	defer fmt.Println("server: hello handler ended")
	fmt.Fprintf(w, "%s", req.Body)
}

func main() {
	fmt.Println("server: listening at \\")

	http.HandleFunc("/", debug)
	http.ListenAndServe(":8090", nil)
}
