package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int = 20

func handle(w http.ResponseWriter, req *http.Request) {

	/*if counter%10 == 0 {
		fmt.Println("server: req", req, "\nheader", req.Header, "\nbody", req.Body)
	}
	counter++*/
	//if _, err := fmt.Fprintf(w, "%s", req.Body); err != nil {
	if _, err := fmt.Fprintf(w, "hello"); err != nil {
		fmt.Println("sending 'hej' error", err)
	}
}

func main() {
	fmt.Println("server: listening at :8090/")
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
