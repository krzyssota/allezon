package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int = 20

func debug(w http.ResponseWriter, req *http.Request) {

	/*if counter%10 == 0 {
		fmt.Println("server: req", req, "\nheader", req.Header, "\nbody", req.Body)
	}
	counter++*/
	//if _, err := fmt.Fprintf(w, "%s", req.Body); err != nil {
	if _, err := fmt.Fprintf(w, "hej"); err != nil {
		fmt.Println("sending 'hej' error", err)
	}
}

func main() {
	fmt.Println("server: listening at \\")

	http.HandleFunc("/", debug)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
